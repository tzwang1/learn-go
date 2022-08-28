package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func is_prime(num int) bool {
	if num == 2 {
		return true
	} else if num%2 == 0 {
		return false
	}
	for i := 3; i < num; i = i + 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func next_prime(cur_prime int) (next int) {
	cur_num := cur_prime + 1
	for {
		if is_prime(cur_num) {
			return cur_num
		}
		cur_num += 1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	find_next_prime := true
	cur_prime := 1
	for find_next_prime {
		fmt.Print("Find next prime? (Y/N)")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Got an error reading input: ", err)
		}
		text = strings.TrimSuffix(text, "\n")
		if text == "Y" {
			cur_prime = next_prime(cur_prime)
			fmt.Println(cur_prime)
		} else if text == "N" {
			find_next_prime = false
		} else {
			log.Println("Invalid input, only accepts 'Y' or 'N' as input.")
		}
	}
}
