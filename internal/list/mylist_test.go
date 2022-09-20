package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleList_Add(t *testing.T) {
	TestsingleList := InitList()
	name := "First name"
	TestsingleList.Add(name + "a")
	TestsingleList.Add(name + "b")
	assert.NotEmpty(t, TestsingleList, "List should not be empty")
	assert.Equal(t, 2, TestsingleList.Size(), "The size should be 3")
}

func TestSingleList_Append(t *testing.T) {
	TestsingleList := InitList()
	name := "First name"
	TestsingleList.Add(name + "a")
	TestsingleList.Add(name + "b")
	assert.NotEmpty(t, TestsingleList, "List should not be empty")
	TestsingleList.Append(name + "c")

	assert.Equal(t, 2, TestsingleList.Size(), "The size should be 3")
}
