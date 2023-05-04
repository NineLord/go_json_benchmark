package main

import (
	// "encoding/json"
	"fmt"
	"github.com/NineLord/go_json_benchmark/pkg/utils/Vector"
	jsoniter "github.com/json-iterator/go"
	"github.com/struCoder/pidusage"
	"os"
	"reflect"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	test13ThreadGetMeTimeForFkSake()
}

func test13ThreadGetMeTimeForFkSake() {
	receiver := make(chan []int64)
	sender := make(chan bool)
	go func(sender chan []int64, receiver chan bool, interval uint) {
		durationInterval := time.Duration(interval)
		result := make([]int64, 0)
		for {
			select {
			case signal := <-receiver:
				fmt.Printf("Thread: received a closing signal! signal: %t\n", signal)
				sender <- result
				close(sender)
				return
			default:
				result = append(result, time.Now().UnixMilli())
				time.Sleep(durationInterval * time.Millisecond)
			}
		}
	}(receiver, sender, 100)
	time.Sleep(500 * time.Millisecond)
	sender <- true
	close(sender)
	result := <-receiver
	for _, res := range result {
		fmt.Println(res)
	}
}

func test12ReadChannel() {
	fmt.Println("Test start!")
	myChannel := make(chan int)
	myChannel <- 1
	myChannel <- 2
	myChannel <- 3
	fmt.Println("Done inserting!")
	for index := 0; index < 3; index++ {
		fmt.Println(<-myChannel)
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			close(c)
			fmt.Println("quit")
			return
		}
	}
}

func test11Fib() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for msg := range c {
			fmt.Println(msg)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func test10ThreadGetTime() {
	receiver := make(chan bool)
	sender := make(chan bool)
	fmt.Println("Channels created!")
	go getTime(receiver, sender, 100)
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Closing thread")
	sender <- true
	close(sender)
	fmt.Println("Starting to collect messages")
	for message := range receiver {
		fmt.Printf("Message: %t\n", message)
	}
	fmt.Println("Main thread has finished!")
}

func getTime(sender chan bool, receiver chan bool, interval uint) {
	// timeToSend := time.Now().UnixMilli()
	durationInterval := time.Duration(interval)
	for {
		fmt.Println("Thread: First line of the loop")
		select {
		// case sender <- timeToSend:
		// 	fmt.Println("Thread: Sent the time!")
		// 	time.Sleep(durationInterval * time.Millisecond)
		// 	timeToSend = time.Now().UnixMilli()
		case signal := <-receiver:
			fmt.Printf("Thread: received a closing signal! signal: %t\n", signal)
			// close(sender)
			return
		default:
			sender <- false
			fmt.Println("Thread: Sent the time!")
			time.Sleep(durationInterval * time.Millisecond)
			// timeToSend = time.Now().UnixMilli()
			// fmt.Println("Thread: HELP ME")
		}

	}
}

func test9ThreadSleep() {
	go printTime(100)
	time.Sleep(500 * time.Millisecond)
}

func printTime(interval uint) {
	durationInterval := time.Duration(interval)
	for {
		fmt.Println(time.Now().UnixMilli())
		time.Sleep(durationInterval * time.Millisecond)
	}
}

func test8PcUsage() {
	sysInfo, err := pidusage.GetStat(os.Getpid())
	if err != nil {
		panic(err)
	}
	fmt.Printf("CPU: %.3f%% \t RAM: %.3fMB\n", sysInfo.CPU, (sysInfo.Memory/1024)/1024)
	fmt.Println(sysInfo)
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
