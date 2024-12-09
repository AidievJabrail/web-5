package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)
	var tmp int
	go func() {
		defer close(out)
		select {
		case tmp = <-firstChan:
			out <- tmp * tmp
			return
		case tmp = <-secondChan:
			out <- tmp * 3
			return

		case <-stopChan:
			return
		}
		
	}()
	
	return out

}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})
	go func(){
		secondChan <- 8
	}()
	fmt.Println(<-calculator(firstChan, secondChan, stopChan))

}
