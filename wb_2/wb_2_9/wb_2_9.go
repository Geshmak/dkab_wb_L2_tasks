package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func foo(fileName string, urlPath string) error {

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res, err := http.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(content)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	args := os.Args[1:]

	fileName := args[0]
	urlPath := args[1]

	if err := foo(fileName, urlPath); err != nil {
		log.Printf("Error: %s", err)
	}
}
