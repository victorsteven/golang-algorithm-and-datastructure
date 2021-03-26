package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSaveMovie(t *testing.T) {

	m := Movie{
		Name:      "The Matrix",
		Year:      2009,
		BoxOffice: 100000,
	}

	SaveMovie(m)

	assert.Equal(t, len(DataStore), 1)
}

func TestSaveMultipleMovie(t *testing.T) {

	m1 := Movie{
		Name:      "The Matrix",
		Year:      2009,
		BoxOffice: 100000,
	}
	m2 := Movie{
		Name:      "Lord of the Rings",
		Year:      2008,
		BoxOffice: 80000,
	}
	m3 := Movie{
		Name:      "The Beatrics",
		Year:      2009,
		BoxOffice: 100000,
	}

	movies := []struct {
		movie Movie
	}{
		{
			m1,
		},
		{
			m2,
		},
		{
			m3,
		},
	}

	chanMovie := make(chan Movie)

	for _, v := range movies {
		go func(mv Movie, ch chan<- Movie) {
			savedMovie := SaveMovie(mv)
			fmt.Println("this is the saved man: ", savedMovie)
			chanMovie <- savedMovie
			time.Sleep(time.Millisecond)
		}(v.movie, chanMovie)
	}

	fmt.Println("this is the chan value: ", chanMovie)

	time.Sleep(time.Millisecond)
	fmt.Println("this is the datastore: ", DataStore)

	assert.Equal(t, len(DataStore), 3)
}

//func TestSaveMultipleMovie2(t *testing.T) {
//
//	//m1 := Movie{
//	//
//	//}
//	//m2 := Movie{
//	//	Name:      "Lord of the Rings",
//	//	Year:      2008,
//	//	BoxOffice: 80000,
//	//}
//	//m3 := Movie{
//	//	Name:      "The Beatrics",
//	//	Year:      2009,
//	//	BoxOffice: 100000,
//	//}
//
//	movies := []struct{
//		Name string
//		Year int
//		BoxOffice int
//	}{
//		{
//			Name:      "The Matrix",
//			Year:      2009,
//			BoxOffice: 100000,
//		},
//		{
//			Name:      "Lord of the Rings",
//			Year:      2008,
//			BoxOffice: 80000,
//		},
//		{
//			Name:      "The Beatrics",
//			Year:      2009,
//			BoxOffice: 100000,
//		},
//	}
//
//	fmt.Println("this is the struct length: ", len(movies))
//
//	chanMovie := make(chan Movie)
//
//	var savedMovies []Movie
//
//	for _, v := range movies {
//		go SaveMovie2(Movie(v), chanMovie)
//	}
//
//	for i := 0; i < len(movies); i++ {
//		fmt.Println("dropping: ", <-chanMovie)
//		savedMovies = append(savedMovies, <-chanMovie)
//	}
//
//	time.Sleep(time.Millisecond)
//	fmt.Println("this is the datastore: ", DataStore)
// 	//assert.Equal(t, len(savedMovies), 3)
//	assert.Equal(t, len(DataStore), 3)
//}
