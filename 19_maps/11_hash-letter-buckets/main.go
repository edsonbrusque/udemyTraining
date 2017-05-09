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
	buckets := make([]int, 255)

	fmt.Println("Loop over the words...")
	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123])
	//fmt.Println("***************")
	//for i := 0; i <= 255; i++ {
	//	if buckets[i] != 0 {
	//		fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	//	}
	//}
}

func hashBucket(word string) int {
	return int(word[0])
}
