package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

func main() {
	var person Person
	b := []byte(`{"id":1,"name":"Gopher","age":5}`)
	err := json.Unmarshal(b, &person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
}
