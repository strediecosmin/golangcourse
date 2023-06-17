package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("First web requests")

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s type of response is %T",response.Body, response)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	content := string(body)

	fmt.Println(content)

}