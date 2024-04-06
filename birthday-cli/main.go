package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please type your birthday by YYYY-MM-DD format!")
	scanner.Scan()
	text := scanner.Text()

	profile, err := NewProfile(text)
	if err != nil {
		log.Fatal("Failed to parse date")
	}

	age := profile.Age()
	fmt.Printf("Your age is %d\n", age)
}
