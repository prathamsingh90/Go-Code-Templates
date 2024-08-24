package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) <= 1 {
		fmt.Println("No Input File provided, read from os.Stdin")
		inputArr := InputReader(os.Stdin)
		fmt.Printf("Input Array is:%v\n", inputArr)
	} else {
		for _, arg := range os.Args[1:] {
			fmt.Printf("Reading Input from %v\n", arg)
			data, err := os.Open(arg)
			if err != nil {
				fmt.Printf("Error reading file:%s\n", arg)
				continue
			}
			inputArr := InputReader(data)
			fmt.Printf("Input Array is:%v\n", inputArr)
			defer data.Close()
		}
	}
	fmt.Printf("Execution Time:%v\n", time.Since(start).Seconds())
}

func InputReader(f *os.File) []int {
	input := bufio.NewScanner(f)
	inputArr := []int{}

	for input.Scan() {
		if input.Text() != "" {
			num, _ := strconv.ParseInt(input.Text(), 10, 64)
			inputArr = append(inputArr, int(num))
		}
	}
	return inputArr
}
