package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	fmt.Println("Getting file...")
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/WordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Reading file's body...")
	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)

	fmt.Println("Printing the file...")
	fmt.Println(str)
}
