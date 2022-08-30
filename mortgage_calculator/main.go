package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_number_input(reader *bufio.Reader) float32 {
	number_text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Got an error when reading input: ", err)
	}
	number_text = strings.TrimSuffix(number_text, "\n")
	number_float, err := strconv.ParseFloat(number_text, 32)
	if err != nil {
		log.Fatalf("Got an error trying to convert %s into a float", number_text)
	}
	return float32(number_float)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the interest rate in percentage (without percent sign) per year: ")
	interest_rate := read_number_input(reader)

	fmt.Print("Enter the total mortgage amount to pay: ")
	total_mortage_amount := read_number_input(reader)

	fmt.Print("Enter the total number of monthly mortage payments: ")
	total_mortage_payments := read_number_input(reader)

	montly_mortgage_payment := (total_mortage_amount / total_mortage_payments) * (interest_rate / 12)
	fmt.Println(montly_mortgage_payment)
}
