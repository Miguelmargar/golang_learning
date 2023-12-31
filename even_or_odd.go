package main

import "fmt"

func newSlice() {

	newSlice := []int{}

	for i := 1; i <= 10; i++ {
		newSlice = append(newSlice, i)
	}

	for _, v := range newSlice {

		if v%2 == 0 {
			fmt.Println("Even")
		}

		if v%2 != 0 {
			fmt.Println("Odd")
		}
	}

}
