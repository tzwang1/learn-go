package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_int_input(reader *bufio.Reader) int {
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Got an error when trying to reader input")
	}
	text = strings.TrimSuffix(text, "\n")
	num, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Could not convert %s to int", text)
	}
	return num
}

func calculate_change(total_cost, amount_paid int) (change_map map[int]int) {
	if total_cost == amount_paid {
		return
	}
	change := []int{25, 10, 5, 1}
	change_idx := 0
	change_map = make(map[int]int)

	diff := amount_paid - total_cost
	for diff > 0 {
		if change[change_idx] > int(diff) {
			change_idx += 1
			continue
		} else {
			diff -= change[change_idx]
			change_map[change[change_idx]] += 1
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the total cost: ")
	total_cost := read_int_input(reader)

	fmt.Print("Enter the amount paid: ")
	amount_paid := read_int_input(reader)

	if amount_paid < total_cost {
		log.Fatal("You did not pay enough!")
	}
	change_map := calculate_change(total_cost, amount_paid)

	if len(change_map) == 0 {
		fmt.Println("You paid the exact amount, you will get zero change!")
	}
	fmt.Println("You total change is: ")
	for change, change_amt := range change_map {
		fmt.Printf("\t%d %d dollar coins\n", change_amt, change)
	}
}
