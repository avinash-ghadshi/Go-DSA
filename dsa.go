package main

import (
	"fmt"

	algo "dsa/algorithms"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	msg, err := algo.LinearSearch(list, 14)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(msg)
	}
}
