package main

import (
	"fmt"

	"github.com/Sanlim/cinema/movie"
	"github.com/Sanlim/cinema/ticket"
)

func init() {}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.Buy(movieName)

}
