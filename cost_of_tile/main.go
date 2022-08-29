package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_input_and_convert_to_int(reader *bufio.Reader) (num int) {
	cost_text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read input, got error: ", err)
	}
	cost_text = strings.TrimSuffix(cost_text, "\n")
	num, err = strconv.Atoi(cost_text)
	if err != nil {
		log.Fatalf("Could not convert %s to int", cost_text)
	}
	return num

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter tile cost: ")
	cost := read_input_and_convert_to_int(reader)

	fmt.Print("Enter width: ")
	width := read_input_and_convert_to_int(reader)

	fmt.Print("Enter height: ")
	height := read_input_and_convert_to_int(reader)

	fmt.Println(width * height * cost)
}
