package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Get the book Moby Dick...")
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Scan the page...")
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	fmt.Println("Set the split function for the scanning operation...")
	scanner.Split(bufio.ScanWords)

	fmt.Println("Create slice to hold counts...")
	buckets := make([]int, 12)

	fmt.Println("Loop over the words...")
	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets)
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
