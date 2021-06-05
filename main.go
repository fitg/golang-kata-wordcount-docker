package main

import (
	"fmt"

	counter "github.com/fitg/golang-kata-wordcount-docker/counter"
)

func main() {
	fmt.Println(counter.WordCount("foo bar"))
}
