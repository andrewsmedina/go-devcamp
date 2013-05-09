package main

import "fmt"

func soma(x, y int, c chan int) {
	c <- x + y
}

func main() {
	c := make(chan int)
	go soma(1, 1, c)
	go soma(2, 2, c)
	x, y := <-c, <-c
	fmt.Println(x+y)
}
