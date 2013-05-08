package main

import (
	"fmt"
	"time"
)

func pegaAgua(p string) {
	fmt.Printf("%s pegou agua\n", p)
}

func bebeAgua(p string) {
	fmt.Printf("%s esta bebendo agua\n", p)
}

func main() {
	pessoas := []string{"andrews", "linus", "fowler"}
	for _, p := range pessoas {
		pegaAgua(p)
		go bebeAgua(p)
	}
	time.Sleep(1)
}

