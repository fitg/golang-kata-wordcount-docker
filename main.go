package main

import (
	"fmt"

	counter "github.com/fitg/golang-kata-wordcount-docker/counter"
)

func main() {
	// test out if the method works
	fmt.Println(counter.WordCount("foo bar"))
}
