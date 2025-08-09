package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/omjogani/bencoding"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a bencoded string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	input = strings.TrimSpace(input)

	decoded, remaining, err := bencoding.DecodeBencode(input)
	if err != nil {
		log.Fatalf("Error decoding: %v", err)
	}

	fmt.Println("Decoded value:", decoded)
	fmt.Println("Remaining string:", remaining)
}
