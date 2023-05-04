package main

import (
	// "encoding/json"
	"fmt"
	"github.com/NineLord/go_json_benchmark/pkg/utils/Vector"

	jsoniter "github.com/json-iterator/go"
	"reflect"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	test7TypeOfNull()
}

func test7TypeOfNull() {
	fmt.Println(reflect.TypeOf([]int{1, 2, 3}).Kind())
	// fmt.Println(reflect.TypeOf(nil).Kind()) // Panic cus gO sMaRt!
	x := reflect.TypeOf(nil)
	fmt.Println(x)
}

func test6WhatTypeIsVector() {
	vec := Vector.NewVector[int](0)
	vec.Push(1)
	vec.Push(2)

	myType := reflect.TypeOf(vec)
	fmt.Println(myType)
}

func test5Marshal() {
	arr := []int{1, 2, 3}
	res, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	println(string(res))
}

func test4Stringify() {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(`{"a": 1, "b": {"c": 2, "d": {"e": 3}}}`), &data); err != nil {
		panic(err)
	}
	// jsonAsString, err := json.MarshalIndent(data, "", "  ")
	jsonAsString, err := json.MarshalToString(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonAsString))
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
