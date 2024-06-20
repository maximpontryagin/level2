package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const outputFile = "output.txt"

func writeHtmlToFile(outputFile string, response *http.Response) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func wget(outputFile, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	err = writeHtmlToFile(outputFile, response)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Println("Ожидался один url адрес")
		os.Exit(1)
	}
	url := args[0]
	fmt.Println(url)
	err := wget(outputFile, url)
	if err != nil {
		log.Println("Ошибка в wget: ", err)
		os.Exit(1)
	}
}
