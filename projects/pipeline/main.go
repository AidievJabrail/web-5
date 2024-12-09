package main

import "fmt"

func removeDuplicates(inputStream, outputStream chan string) {
	prev := <-inputStream
	outputStream <- prev
	for i := range inputStream {
		if prev != i {
			outputStream <- i
		}
		prev = i
	}
	close(outputStream)

}

func main() {

	input := make(chan string)
	output := make(chan string)
	go removeDuplicates(input, output)
	go func() {
		var a, b, c, d, e string
		fmt.Scan(&a, &b, &c, &d, &e)
		input <- a
		input <- b
		input <- c
		input <- d
		input <- e
		close(input)
	}()
	for i := range output {
		fmt.Println(i)
	}

}
