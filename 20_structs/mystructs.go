package main

import "fmt"

type User struct { // init of struct
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("this file shows the working of structs.")
	// no inheritance in golang no super or parent
	aadi := User{"Aadi", "aadihehe@mail", true, 22}
	fmt.Println(aadi)
	// output : {Aadi aadihehe@mail true 22}
	fmt.Printf("Aadi's details are: %+v\n", aadi)
	// output : Aadi's details are: {Name:Aadi Email:aadihehe@mail Status:true Age:22}
	fmt.Printf("Aadi's details are: name %v email is %v ", aadi.Name, aadi.Email)

}
