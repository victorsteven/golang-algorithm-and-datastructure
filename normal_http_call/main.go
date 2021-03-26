package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {

	http.HandleFunc("/create", Create)
	http.HandleFunc("/create2", Create2())

	fmt.Println("We have the server running")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

type Movie struct {
	Name      string `json:"name"`
	Year      int    `json:"year"`
	BoxOffice int    `json:"box_office"`
}

var DataStore = map[int]Movie{}

var id = 1

func SaveMovie(m Movie) Movie {
	var mutex = &sync.Mutex{}

	mutex.Lock()
	DataStore[id] = m
	id++
	mutex.Unlock()

	return m
}

func SaveMovie2(m Movie, mc chan Movie) {
	var mutex = &sync.Mutex{}

	mutex.Lock()
	DataStore[id] = m
	id++
	mutex.Unlock()

	mc <- m
}

func Getmovie(id int) Movie {
	return DataStore[id]
}

func Create(w http.ResponseWriter, r *http.Request) {

	m := Movie{}
	if r.Method != "POST" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "%s", "This request require a post http verb")
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	err = json.Unmarshal(body, &m)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	result := SaveMovie(m)

	w.WriteHeader(200)
	fmt.Fprintf(w, "%v", result)
}


func Create2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := Movie{}
		if r.Method != "POST" {
			w.WriteHeader(400)
			fmt.Fprintf(w, "%s", "This request require a post http verb")
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

		err = json.Unmarshal(body, &m)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

		result := SaveMovie(m)

		w.WriteHeader(200)
		fmt.Fprintf(w, "%v", result)
	}
}