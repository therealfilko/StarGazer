package main

import (
	"os"

	"github.com/google/go-github/v60/github"
)

func main()  {
    startDate := os.Args[1]
    endDate := os.Args[2]

    search, err := github.SearchService{}
}
