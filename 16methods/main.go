package main

import "fmt"

type repeatStructs struct {
	Name string
	Age int
	Country string
}

type User struct {
	Name string
	Age int
	Country string
	Status bool
	oneAge int
}

func (u User) GetStatus()  {
	fmt.Println("Is user active: ", u.Status)
}

type differentStruct struct {
	name string
}

func main()  {
	fmt.Println("Welcome to methods lesson")

	newVariable := repeatStructs{
		"Cosmin", 31, "Romania",
	}
	fmt.Println(newVariable)

	anotheStruct := differentStruct{"Stredie"}

	fmt.Printf("just some details %+v\n", anotheStruct)
	fmt.Printf("The name is %+v\n", anotheStruct.name)

	newUser := User{"Cosmin",31,"Romania",true,31}

	newUser.GetStatus()
	newUser.NewCountry()
	fmt.Println(newUser)
}

func (u User) NewCountry()  {
	u.Country = "Ukraine"

	fmt.Println("another country is:", u.Country)
}