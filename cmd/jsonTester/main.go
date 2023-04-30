package main

import (
	. "fmt"
)

func main() {
	Println("Ok!")
	a := make([]int, 3, 5)
	// a := []int{1,2,3}
	b := append(a, 8)
	b2 := append(a, 8)
	c := append(b, 8)
	Println(a, len(a), cap(a), "\n",
		a[:cap(a)], "\n",
		b, "\n",
		b2, "\n",
		c)
	Println("Bye!")
}
