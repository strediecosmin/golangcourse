package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main()  {
	fmt.Println("json lesson")

	// EncodeJson()
	decodeJson()
}

func EncodeJson()  {	
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev","js"}},
		{"Golang Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"fullstack","go"}},
		{"Mern Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// package data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
	fmt.Printf("The type is %T\n", finalJson)

}

func decodeJson()  {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev","js"]
		}
		`)

		var lcoCourse course

		checkValidJson := json.Valid(jsonDataFromWeb)

		if checkValidJson {
			fmt.Printf("Json was valid\n")
			json.Unmarshal(jsonDataFromWeb, &lcoCourse)
			fmt.Printf("%#v\n",lcoCourse)
		} else {
			fmt.Println("json was not valid")
		}

		// some cases for key value pairs

		var myOnlineData map[string]interface{}
		json.Unmarshal(jsonDataFromWeb, &myOnlineData)
		fmt.Printf("%#v\n", myOnlineData)
		for k,v := range myOnlineData {
			fmt.Printf("Key is: %v Value is: %v and type is: %T\n", k, v, v)
		}

}