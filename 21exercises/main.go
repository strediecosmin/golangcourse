package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	fmt.Println("Web Verb video - LCO")

	// PerformGetRequest()
	// PerformJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest()  {
	const myUrl = "http://localhost:8000/get"

	resp, err := http.Get(myUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Content length is: ", resp.ContentLength)

	var responseString strings.Builder

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	byteCount, err := responseString.Write(body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(body))

}

func PerformJsonRequest()  {
	const myUrl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with GOLANG",
			"price": 0
		}
	`)

	resp, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}

func PerformPostFormRequest()  {
	const myUrl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}

	data.Add("firstname", "Cosmin")
	data.Add("lastname", "Stredie")

	resp, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()	
	content, _ := io.ReadAll(resp.Body)

	fmt.Println(string(content))

}