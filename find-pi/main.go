package main

// Find PI to the Nth Digit - Enter a number and have the program generate PI up
// to that many decimal places. Keep a limit to how far the program will go.

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const PI = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of digits of pi to show: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Got an invalid input: ", text)
	}
	text = strings.TrimSuffix(text, "\n")
	digits, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Could not convert %s to an int.", text)
	}
	if digits > 100 {
		log.Fatalf("Maximum number of digits that can be displayed is 100.")
	}
	var pi_truncated string
	if digits == 1 {
		pi_truncated = "3"
	} else {
		pi_truncated = "3." + PI[2:digits+1]
	}
	fmt.Printf("The first %d digits of pi are %s", digits, pi_truncated)
}
