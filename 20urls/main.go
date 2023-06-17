package main

import (
	"fmt"
	"log"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gibrrish"

func main()  {
	fmt.Println("How to handle urls")

	fmt.Println(myUrl)

	// parse the URL
	result, err := url.Parse(myUrl)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of query is %T\n", qparams)

	for _, v := range qparams {
		fmt.Println(v)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()

	fmt.Println(anotherURL)

}