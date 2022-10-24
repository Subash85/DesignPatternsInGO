package main

import (
	"fmt"
)
import "LinkedListGo/internal/list"

// add , insert and remove  . add unit test file. Add [name1, Name2]

func main() {
	singleList := list.List()
	fmt.Printf("Add: A\n")
	singleList.Add("A")
	fmt.Printf("Add: B\n")
	singleList.Add("B")
	fmt.Printf("Add: C\n")
	singleList.Add("C")
	println("list before add \n")
	singleList.Print()

	singleList.InsertAt(2, "V")

	println("list after add\n")
	singleList.Print()

	_, err := singleList.Insert("X", "Y", "", "Q")
	if err != nil {
		println(err.Error())
	}
	println("list after Multiple inserts")

	singleList.Print()

	fmt.Printf("\n Size: %d", singleList.Size())

}
