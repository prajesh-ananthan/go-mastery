package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	f, _ := os.Open("data.txt")

	// Create a Scanner for the file
	scanner := bufio.NewScanner(f)

	// Read and print each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
