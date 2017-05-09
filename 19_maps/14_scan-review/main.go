package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

func main() {
	// An artificial input source.
	const input = "Este é um teste de entrada de dados."

	scanner := bufio.NewScanner(strings.NewReader(input))

	//Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Count the words.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err:= scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading input:", err)
	}
}
