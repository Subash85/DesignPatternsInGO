package list

import (
	"errors"
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
	tail *ele
}

func List() *singleList {
	return &singleList{}
}

func (s *singleList) Add(name string) {
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

func (s *singleList) Insert(addStrings ...string) (res string, err error) {

	if len(addStrings) == 0 {
		return "", errors.New("No string to insert")
	}
	var containsErr error = nil
	for _, word := range addStrings {
		if !checkEmptyString(word) {
			s.InsertAt(s.Size(), word)
		} else {
			containsErr = errors.New("Empty strings not inserted")
		}
	}
	if containsErr != nil {
		return "", containsErr
	}
	return "", nil
}

func checkEmptyString(checkEmpty string) bool {
	strings.TrimSpace(checkEmpty)
	return len(checkEmpty) == 0
}

func (s *singleList) Get(index int) (returnString string, err error) {
	var matchedString string
	if s.Size() < index {
		return "", fmt.Errorf("Array Out of bound")
	}
	current := s.head
	for i := 0; i < s.len; i++ {
		if i == index {
			matchedString = current.name
			break
		}
		current = current.next
	}
	return matchedString, nil
}

func (s *singleList) InsertAt(index int, name string) error {
	if s.len < 0 {
		return fmt.Errorf("removeBack: List is empty")
	}
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
		s.head.next = nil
		s.tail = ele
		s.tail.next = nil
		s.len++
	} else if index == 0 {
		s.head = ele
		s.head.next = s.tail
		s.len++
	} else if index == s.len {
		s.tail = ele
		s.loopSwap(index, ele)
	} else {
		s.loopSwap(index, ele)
	}
	return nil
}

func (s *singleList) loopSwap(index int, eleSwap *ele) *ele {
	current := s.head
	for i := 0; i < s.len; i++ {

		if i == index-1 && index > 0 {
			swap := current.next
			current.next = eleSwap
			eleSwap.next = swap
			s.len++
		}
		current = current.next
	}
	return current
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

func (s *singleList) Print() {
	next := s.head
	print("[")
	for i := 1; i <= s.len; i++ {
		print(next.name)
		if i != s.len {
			print(",")
		}
		next = next.next
	}
	print("]")
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
