package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Getting file...")
	res, err := http.Get("http://www-01_value.sil.org/linguistics/wordlists/english/wordlist/WordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Reading file's body...")
	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)

	fmt.Println("Printing the file...")
	fmt.Println(str)
}
