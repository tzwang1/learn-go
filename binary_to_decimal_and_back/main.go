package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func reverse_string(str string) string {
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func convert_binary_to_decimal(binary string) int {
	decimal := 0
	idx := 0
	for i := len(binary) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(binary[i]))
		if err != nil {
			log.Fatalf("Could not convert %s to int", string(binary[i]))
		}
		decimal = decimal + (digit * int(math.Pow(2, float64(idx))))
		idx++
	}
	return decimal
}

func convert_decimal_to_binary(decimal int) string {
	binary_str := ""
	for decimal > 0 {
		binary_str += strconv.FormatInt(int64(decimal%2), 10)
		decimal = decimal / 2
	}
	return reverse_string(binary_str)
}

func validate_binary(text string) error {
	for i := 0; i < len(text); i++ {
		if text[i] != '1' && text[i] != '0' {
			return fmt.Errorf("got an invalid argument: %s is not a binary number", string(text[i]))
		}
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Convert decimal or binary? (D/B)")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Got error when reading input: ", err)
	}
	text = strings.TrimSuffix(text, "\n")
	if text == "D" {
		fmt.Print("Enter decimal number to convert: ")
		decimal_str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Got an error when trying to read input: ", err)
		}
		decimal_str = strings.TrimSuffix(decimal_str, "\n")
		decimal, err := strconv.Atoi(decimal_str)
		if err != nil {
			log.Fatalf("Could not convert %s to int", decimal_str)
		}
		fmt.Println(convert_decimal_to_binary(decimal))
	} else if text == "B" {
		fmt.Print("Enter binary number to convert: ")
		binary, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Got an error when trying to read input: ", err)
		}
		binary = strings.TrimSuffix(binary, "\n")
		err = validate_binary(binary)
		if err != nil {
			log.Fatal("Got an invalid binary number as input: ", err)
		}
		fmt.Println(convert_binary_to_decimal(binary))
	} else {
		log.Fatalf("Got an invalid input %s, can only accept 'D' or 'B'", text)
	}
}
