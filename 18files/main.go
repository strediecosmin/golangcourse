package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fmt.Println("files lesson after vacation")

	content := "write something random in a file"

	file, err := os.Create("./randomfile.txt")

	if err != nil {
		panic(err)
	}

	length, err :=	io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./randomfile.txt")
}

func readFile(fileName string)  {

	databyte, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Text data is in the file \n", string(databyte))
	os.Stdout.Write(databyte)
}