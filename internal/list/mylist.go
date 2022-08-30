package list

import (
	"fmt"
	"strings"
)

type ele struct {
	name string
	next *ele
}

type singleList struct {
	len  int
	head *ele
}

func InitList() *singleList {
	return &singleList{}
}

func (s *singleList) Add(name string) {
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
	} else {
		ele.next = s.head
		s.head = ele
	}
	s.len++
	return
}

func (s *singleList) Append(name string) {
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = ele
	}
	s.len++
	return
}
func (s *singleList) InsertAt(index int, name string) error {
	//singleList := s
	if s.len < 0{
		return fmt.Errorf("removeBack: List is empty")
	}
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
		s.len++
	} else if index == 0 {
		current := s.head
		s.head=ele
		s.head.next = current
		current = s.head
		s.len++
	} else if index == s.len {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = ele
		s.len++
	} else {
		current := s.head
		for i := 0; i < s.len; i++ {

			if i == index-1 && index > 0 {
				swap := current.next
				current.next = ele
				ele.next = swap
				s.len++
			}
			current = current.next
		}
	}
	return nil
}
func (s *singleList) RemoveFront() error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}

func (s *singleList) RemoveBack() error {
	if s.head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}
	var prev *ele
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
	return nil
}

func (s *singleList) Front() (string, error) {
	if s.head == nil {
		return "", fmt.Errorf("Single List is empty")
	}
	return s.head.name, nil
}

func (s *singleList) Size() int {
	return s.len
}
func (s *singleList) PrintAll()  string {
	next := s.head

	var listToString  = "["

	for i := 1; i <= s.len; i++ {

		listToString += next.name
		if i != s.len {
			listToString +=","
		}
		next = next.next
	}
	listToString += "]"

	return strings.TrimSpace(listToString)
}

func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}