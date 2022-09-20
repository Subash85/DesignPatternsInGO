package main

import (
	"fmt"
)
import "LinkedListGo/internal/list"

// add , insert and remove  . add unit test file. Add [name1, Name2]

func main() {
	singleList := list.InitList()
	fmt.Printf("Add: A\n")
	singleList.Add("A")
	fmt.Printf("Add: B\n")
	singleList.Add("B")
	fmt.Printf("Add: C\n")
	singleList.Append("C")
	fmt.Printf(" list before add %v \n", singleList.PrintAll())

	fmt.Println("insert : v")

	singleList.InsertAt(2, "V")

	fmt.Printf(" list after  add %v \n", singleList.PrintAll())

	fmt.Printf("Size: %d \n", singleList.Size())

	//err := singleList.Traverse()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	/*	fmt.Printf("RemoveFront\n")
		err := singleList.RemoveFront()
		if err != nil {
			fmt.Printf("RemoveFront Error: %s\n", err.Error())
		}

		fmt.Printf("RemoveBack\n")
		err = singleList.RemoveBack()
		if err != nil {
			fmt.Printf("RemoveBack Error: %s\n", err.Error())
		}

		fmt.Printf("RemoveBack\n")
		err = singleList.RemoveBack()
		if err != nil {
			fmt.Printf("RemoveBack Error: %s\n", err.Error())
		}

		fmt.Printf("RemoveBack\n")
		err = singleList.RemoveBack()
		if err != nil {
			fmt.Printf("RemoveBack Error: %s\n", err.Error())
		}

		err = singleList.Traverse()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("Size: %d\n", singleList.Size())*/
}
