package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := [...]int{1, 2, 3}

	fmt.Printf("%d,%d\n", len(x), cap(x))
	fmt.Printf("%d,%d\n", len(y), cap(y))
	sli := make([]int, 0, 3)
	fmt.Println(sli)

}
