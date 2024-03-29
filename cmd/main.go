package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v60/github"
	"github.com/joho/godotenv"
)

func main()  {
    token := getToken()
    client := github.NewClient(nil).WithAuthToken(token)

    startDate := os.Args[1]
    endDate := os.Args[2]

    query := "created:" + startDate + ".." + endDate

    fmt.Printf("Searchquery: %s\n", query)
    fmt.Printf("Anfangsdatum: %s\n", startDate)
    fmt.Printf("Enddatum: %s\n", endDate)

    searchRepositories(client, query)
}

func getToken() string  { 
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    return os.Getenv("AUTH_TOKEN")
}

func searchRepositories(client *github.Client, query string) {
    ctx := context.Background()

    opts := &github.SearchOptions{Sort: "stars", Order: "desc"}
    result, _, err := client.Search.Repositories(ctx, query, opts)
    if err != nil {
        log.Fatal("Couldn't find it")
    }

    fmt.Printf("Gefunden Repos: %d\n", *result.Total)
    for _, repo := range result.Repositories {
        fmt.Printf("Name: %s, Sterne: %d, URL: %s\n", *repo.Name, *repo.StargazersCount, *repo.HTMLURL)
    }
}

func testAuthentication(client *github.Client) {
    user, _, err := client.Users.Get(context.Background(), "")
    if err != nil {
        log.Fatalf("Authentifizierungsfehler: %v", err)
    } else {
        fmt.Printf("Erfolgreich authentifiziert als Benutzer: %s\n", *user.Login)
    }
}
