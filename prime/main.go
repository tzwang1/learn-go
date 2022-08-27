package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func find_primes(num int) (primes []int) {
	c := 2
	for num > 1 {
		if num%c == 0 {
			primes = append(primes, c)
			num = num / c
		} else {
			c++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number to find prime factors for: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Got an error when reading input: ", err)
	}
	text = strings.TrimSuffix(text, "\n")
	num, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Could not convert input: %s into an int", text)
	}
	fmt.Println(find_primes(num))
}
