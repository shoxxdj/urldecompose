package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	// Trim the newline character
	input = strings.TrimSpace(input)

	// Split input into URL and arguments
	parts := strings.Split(input, "?")
	fmt.Println(parts[0])

	if len(parts) > 1 {
		args := strings.Split(parts[1], "&")
		for _, arg := range args {
			fmt.Println(arg)
		}
	} else {
		fmt.Println("No arguments provided.")
	}
}

