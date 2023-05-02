package main

import (
	"fmt"
	"reflect"
)

func main() {
	test3AddingToArray()
}

func test3AddingToArray() {
	arrayies := make([]int, 0)
	arrayies = addItSlice(arrayies)
	arrayies = addItSlice(arrayies)
	arrayies = addItSlice(arrayies)
	arrayies = addItSlice(arrayies)
	arrayies = addItSlice(arrayies)
	arrayies = addItSlice(arrayies)
	fmt.Println(arrayies, len(arrayies), cap(arrayies))
}

func addItSlice(mySlice []int) []int {
	return append(mySlice, 5)
}

func test2AddingToMap() {
	mapy := make(map[string]interface{})
	addIt(mapy)
	fmt.Println(mapy)
}

func addIt(myMap map[string]interface{}) {
	myMap["a"] = 5
}

func test1() {
	// x := make([]int, 0, 3)
	z := [3]int{0, 0, 0}
	x := z[:0]
	y := append(x, 1)
	y = append(x, 2)
	fmt.Println("z:", z)
	innerMain(x)
	fmt.Println(y)
}

func innerMain(x interface{}) {
	switch reflect.TypeOf(x).Kind() {
	case reflect.Array:
		doSomethingWithArray(x.([3]int))
	case reflect.Slice:
		doSomethingWithSlice(x.([]int))
	default:
		panic("Wrong type!")
	}
}

func doSomethingWithSlice(arr []int) {
	fmt.Println("Slice", arr, len(arr), cap(arr))
}

func doSomethingWithArray(arr [3]int) {
	fmt.Println("Array", arr)
}
