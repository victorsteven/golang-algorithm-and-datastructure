package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
func main(){
	cus := Customer{
		Name: "Mike",
		Age:  30,
	}
	str, err := struct_to_json(cus)
	if err != nil {
		fmt.Printf("the input must be of type Customer struct, but '%T' was given\n", cus)
	}
	fmt.Println(str)

	jsonValue := `{"name": "oke", "age": 60}`

	the_struct, err := json_to_struct(jsonValue)
	if err != nil {
		fmt.Println("Cannot convert json to struct")
	}
	fmt.Println(the_struct)
}

func struct_to_json(val interface{}) (string, error) {
	theByte, err := json.Marshal(val.(Customer)) //make sure the interface is a Customer struct
	if err != nil {
		return "", err
	}
	return string(theByte), nil
}

func json_to_struct(str string) (Customer, error) {
	the_struct := Customer{}
	err := json.Unmarshal([]byte(str), &the_struct)
	if err != nil {
		return the_struct, err
	}
	return the_struct, nil
}
