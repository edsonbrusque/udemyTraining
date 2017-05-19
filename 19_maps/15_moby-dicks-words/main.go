package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Get the Moby Dick book...")
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Scan the page...")
	// NewScanner takes a reader and res.Body implements the reader interface
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	fmt.Println("Set the split function for the scanning operation...")
	scanner.Split(bufio.ScanWords)

	fmt.Println("Create a slice of slice of string to hold slices of words...")
	buckets := make([][]string, 12)

	fmt.Println("Create slices to hold words...")
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	fmt.Println("Loop over the words...")
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	fmt.Println("Print len of each bucket...")
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	fmt.Println("Print the words in one of the buckets...")
	fmt.Println(buckets[6])
}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets

	// A more uneven distribution
	//return len(word) % buckets
}
