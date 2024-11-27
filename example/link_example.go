package main

import (
	"fmt"
	"persistlink/lib" 
)

func main() {

	list := lib.Empty[int]().Append(10).Append(20).Append(30)
	anotherList := lib.Empty[int]().Append(40).Append(50)
	concatenatedList := list.Concat(anotherList)


	fmt.Println("Concatenated list as slice:", concatenatedList.AsSlice()) // output is [30, 20, 10, 50, 40]
}
