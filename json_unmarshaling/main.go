package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Status int `json:"status"`
}

func main() {
	var data = []byte(`{"status": 200}`)
	var result = Result{}

	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("result=%+v\n", result)
}





