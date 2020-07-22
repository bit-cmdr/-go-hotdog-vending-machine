package main

import (
	"fmt"
	"log"
	"strconv"
)

func hotdogMachine(inventoryCount int) (chan<- string, <-chan string) {
	in := make(chan string)
	out := make(chan string)
	go func(hc int, in, out chan string) {
		for {
			currency := <-in
			switch {
			case hc > 0:
				switch {
				case currency == "dollar":
					hc--
					out <- "hotdog"
				default:
					out <- "wilted lettuce"
				}
			default:
				out <- "all out"
			}
		}
	}(inventoryCount, in, out)
	return in, out
}

func main() {
	var input string

	fmt.Print("How many hotdogs in the vending machine? ")
	fmt.Scanln(&input)
	count, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Unable to convert number of hot dogs for vending machine. Expected a number, received", input)
	}

	fmt.Print("How many hotdogs would you like? ")
	fmt.Scanln(&input)
	req, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Unable to convert number of hot dogs wanted. Expected a number, received", input)
	}

	in, out := hotdogMachine(count)
	defer close(in)
	go func(out <-chan string) {
		for {
			fmt.Println(<-out)
		}
	}(out)
	in <- "pocket lint"
	for i := 0; i < req; i++ {
		in <- "dollar"
	}

	fmt.Scanln(&input)
}
