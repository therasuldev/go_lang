//! channel
package main

import (
	"fmt"
)

func main() {
	mychannel := make(chan string)
	go mychan(mychannel)
	fmt.Println(<-mychannel)
	pointValue := 10
	val := myStrc{"go dilinde interface kavrami"}
	fmt.Println(pointer(&pointValue))
	golang(val)
	//!--------------------- Map , Array , switch ----------
	myArr()
	fmt.Println()
	mySwitch(0)
	myMap()
}

func mychan(chan1 chan string) {
	chan1 <- "go derslerinin temelini bitirdik"
}

//? pointer
func pointer(a *int) int {
	return *a * 10
}

//?interface and struct
type myStrc struct {
	I string
}
type myIntr interface {
	myFunc() string
}

func (types myStrc) myFunc() string {
	return types.I
}
func golang(son myIntr) {
	fmt.Println(son.myFunc())
}

//? /////////////////
func myArr() {
	array := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i += 1 {
		fmt.Print(array[i])
	}
}
func mySwitch(value int) {
	switch value {
	case 0:
		fmt.Println("Switch metodu ")
	default:
		fmt.Println("default value")
	}
}
func myMap() {
	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for key, val := range myMap {
		fmt.Printf("{key:%v, val:%v} ", key, val)
	}

}
