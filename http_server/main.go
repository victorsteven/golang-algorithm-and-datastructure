package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello\n")
	})
	log.Fatal(http.ListenAndServe(":8080", mux))

	http.HandleFunc("/me", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "Hello")
	})
}

//something else
//func ore() {
//
//	age := 27
//	rows, err := db.Query("SELECT name FROM users WHERE age=?", age)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close() // A
//	for rows.Next() { //B
//		var name string
//		if err := rows.Scan(&name); err != nil { //C
//			log.Fatal(err)
//		}
//		fmt.Printf("%s is %d\n", name, age)
//	}
//	if err := rows.Err(); err != nil { //D
//		log.Fatal(err)
//	}
//}
