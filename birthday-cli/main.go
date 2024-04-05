package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please type input text!")
	scanner.Scan()
	text := scanner.Text()

	profile, err := NewProfile(text)
	if err != nil {
		log.Fatal("Failed to parse date")
	}

	fmt.Printf("Birthday is %s\n", profile.Birthday.String())
}
