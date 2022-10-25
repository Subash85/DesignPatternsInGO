package list

import (
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
	singleListTest.Add("A")
	singleListTest.Add("B")
	singleListTest.Add("C")
	singleListTest.InsertAt(insertindex, "V")
	for i := 0; i < singleListTest.Size(); i++ {
		if i == insertindex {
			actual, err := singleListTest.Get(insertindex)
			assert.Nil(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, "V", actual)
		}
	}
}

func TestSingleList_InsertAtLastIndex(t *testing.T) {
	singleListTest := List()
	var insertindex = 3
	singleListTest.Add("A")
	singleListTest.Add("B")
	singleListTest.Add("C")
	singleListTest.InsertAt(insertindex, "V")
	for i := 0; i < singleListTest.Size(); i++ {
		if i == insertindex {
			actual, err := singleListTest.Get(insertindex)
			assert.Nil(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, "V", actual)
		}
	}
}

func TestSingleList_InsertAtFirst(t *testing.T) {
	singleListTest := List()
	var insertindex = 0
	singleListTest.InsertAt(insertindex, "V")
	for i := 0; i < singleListTest.Size(); i++ {
		if i == insertindex {
			actual, err := singleListTest.Get(insertindex)
			assert.Nil(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, "V", actual)
		}
	}
}

func TestSingleList_Insert(t *testing.T) {

	singleListTest := List()
	singleListTest.Add("A")
	response, err := singleListTest.Insert("a", "b", "C")
	assert.Nil(t, err)
	assert.Equal(t, "", response)

}
func TestSingleList_Insert_blanks(t *testing.T) {

	singleListTest := List()
	response, err := singleListTest.Insert("", "A", "A B", " ABC")
	assert.NotNil(t, err)
	assert.Equal(t, "", response)
	assert.Equal(t, "Empty strings not inserted", err.Error())
	assert.Equal(t, 3, singleListTest.Size())

}

func TestSingleList_Get(t *testing.T) {

	singleListTest := List()
	response, err := singleListTest.Insert("a", "b", "C")
	assert.Nil(t, err)
	assert.Equal(t, "", response)
	singleListTest.Print()
	retriveString, err := singleListTest.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, retriveString)
	assert.Equal(t, "C", retriveString)

}

func TestSingleList_Get_OutOfBound(t *testing.T) {
	singleListTest := List()
	_, err := singleListTest.Get(2)
	assert.NotNil(t, err)
	assert.Equal(t, "Array Out of bound", err.Error())

}
