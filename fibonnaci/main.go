package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func generate_fib(n int) []int {
	fib_nums := []int{1, 1}
	if n == 1 {
		return []int{1}
	} else if n == 2 {
		return []int{1, 1}
	}
	i := 3
	for i <= n {
		fib_nums = append(fib_nums, fib_nums[len(fib_nums)-1]+fib_nums[len(fib_nums)-2])
		i++
	}
	return fib_nums
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of Fibonnaci numbers to display: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Got an error while reading input: ", err)
	}
	text = strings.TrimSuffix(text, "\n")
	num, err := strconv.Atoi(text)
	if err != nil {
		log.Fatalf("Could not convert %s to an int", text)
	}
	fmt.Println(generate_fib(num))
}
