package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v60/github"
	"github.com/joho/godotenv"
)

// 1. ich muss env richtig einbinden also mein Token | DONE
// 2. danach authentifizieren | DONE
// 3. Suche definieren
// 4. testen

func main()  {
    token := getToken()
    client := github.NewClient(nil).WithAuthToken(token)
    testAuthentication(client)

    startDate := os.Args[1]
    endDate := os.Args[2]

    fmt.Printf("Anfangsdatum: %s\n", startDate)
    fmt.Printf("Enddatum: %s\n", endDate)
}

func getToken() string  {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    return os.Getenv("AUTH_TOKEN")
}

func testAuthentication(client *github.Client) {
    user, _, err := client.Users.Get(context.Background(), "")
    if err != nil {
        log.Fatalf("Authentifizierungsfehler: %v", err)
    } else {
        fmt.Printf("Erfolgreich authentifiziert als Benutzer: %s\n", *user.Login)
    }
}
