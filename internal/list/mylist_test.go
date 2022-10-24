package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleList_Add(t *testing.T) {
	TestsingleList := List()
	name := "First name"
	TestsingleList.Add(name + "a")
	TestsingleList.Add(name + "b")
	assert.NotEmpty(t, TestsingleList, "List should not be empty")
	assert.Equal(t, 2, TestsingleList.Size(), "The size should be 3")
}

func TestSingleList_Append(t *testing.T) {
	TestsingleList := List()
	name := "First name"
	TestsingleList.Add(name + "a")
	TestsingleList.Add(name + "b")
	assert.NotEmpty(t, TestsingleList, "List should not be empty")
	TestsingleList.Add(name + "c")

	assert.Equal(t, 3, TestsingleList.Size(), "The size should be 3")
}

func TestSingleList_InsertAt(t *testing.T) {
	singleListTest := List()
	var insertindex = 2

	fmt.Printf("Add: A\n")
	singleListTest.Add("A")
	fmt.Printf("Add: B\n")
	singleListTest.Add("B")
	fmt.Printf("Add: C\n")
	singleListTest.Add("C")
	fmt.Printf(" list before add \n")
	singleListTest.Print()
	fmt.Println("inserting here : v")

	singleListTest.InsertAt(insertindex, "V")
	fmt.Printf(" list after add  \n")
	singleListTest.Print()

	for i := 0; i < singleListTest.Size(); i++ {
		if i == insertindex {
			///assert.Equal(t, "V", singleListTest[i])
		}

	}

}

func TestSingleList_Insert(t *testing.T) {

	singleListTest := List()
	singleListTest.Add("A")
	response, err := singleListTest.Insert("a", "b", "C")
	assert.Nil(t, err)
	assert.Nil(t, response)

}
func TestSingleList_Insert_blanks(t *testing.T) {

	singleListTest := List()
	response, err := singleListTest.Insert("", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, "Empty strings not inserted", err.Error())
	assert.Nil(t, response)

}
