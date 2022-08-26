package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const E = "2.7182818284590452353602874713526624977572470936999595749669676277240766303535945713821785251664274274"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of decimal places of e to return: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Got an invalid input trying to read: ", text)
	}
	text = strings.TrimSuffix(text, "\n")
	digits, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("Got an input that could not be converted to an int: ", text)
	}
	if digits > 100 {
		log.Fatal("Can only display e with at most 100 decimal points.")
	}
	fmt.Printf("e with %d decimals is: %s", digits, E[:2+digits])
}
