package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Getting file...")
	res, err := http.Get("http://www-01_value.sil.org/linguistics/wordlists/english/wordlist/WordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Making a map...")
	words := make(map[string]string)

	fmt.Println("Scanning through the file's body...")
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	fmt.Println("Writing contents to table...")
	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Println("Printing...")
	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}
