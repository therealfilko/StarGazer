package main

import (
	"fmt"
	"os"
)

// 1. ich muss env richtig einbinden also mein Token
// 2. danach authentifizieren
// 3. Struct definieren
// 4. Suche definieren
// 5. testen
func main()  {
    startDate := os.Args[1]
    endDate := os.Args[2]
    token := os.Getenv("AUTH_TOKEN")

    fmt.Println(startDate)
    fmt.Println(endDate)
    fmt.Println(token)

}
