package stack

import "testing"

func TestLen_String_Empty(t *testing.T) {
	s := Stack[string]{}
	len := s.Len()
	if len != 0 {
		t.Fail()
	}
}

func TestLen_String_One(t *testing.T) {
	s := Stack[string]{}
	s.Push("Hello")
	len := s.Len()
	if len != 1 {
		t.Fail()
	}
}

func TestLen_Float_Empty(t *testing.T) {
	s := Stack[float64]{}
	len := s.Len()
	if len != 0 {
		t.Fail()
	}
}

func TestLen_Float_One(t *testing.T) {
	s := Stack[float64]{}
	s.Push(1.1)
	len := s.Len()
	if len != 1 {
		t.Fail()
	}
}

func TestPush_String(t *testing.T) {
	s := Stack[string]{}
	s.Push("Hello")
	s.Push("World")
	if s.head.data != "Hello" {
		t.Fail()
	}
	if s.head.nextItem.data != "World" {
		t.Fail()
	}
}

func TestPush_Float(t *testing.T) {
	s := Stack[float64]{}
	s.Push(1.1)
	s.Push(2.2)
	if s.head.data != 1.1 {
		t.Fail()
	}
	if s.head.nextItem.data != 2.2 {
		t.Fail()
	}
}

func TestPop_String(t *testing.T) {
	s := Stack[string]{}
	s.Push("Hello")
	s.Push("World")
	s.Pop()
	if s.head.data != "Hello" {
		t.Fail()
	}
	if s.head.nextItem != nil {
		t.Fail()
	}
}

func TestPop_Float(t *testing.T) {
	s := Stack[float64]{}
	s.Push(1.1)
	s.Push(2.2)
	s.Pop()
	if s.head.data != 1.1 {
		t.Fail()
	}
	if s.head.nextItem != nil {
		t.Fail()
	}
}

func TestGetByIndex_String(t *testing.T) {
	s := Stack[string]{}
	s.Push("Hello")
	s.Push("World")
	s.Push("!")
	if s.GetByIndex(0).data != "Hello" {
		t.Fail()
	}
	if s.GetByIndex(1).data != "World" {
		t.Fail()
	}
	if s.GetByIndex(2).data != "!" {
		t.Fail()
	}
}
func TestGetByIndex_Float(t *testing.T) {
	s := Stack[float64]{}
	s.Push(1.1)
	s.Push(2.2)
	s.Push(3.3)
	if s.GetByIndex(0).data != 1.1 {
		t.Fail()
	}
	if s.GetByIndex(1).data != 2.2 {
		t.Fail()
	}
	if s.GetByIndex(2).data != 3.3 {
		t.Fail()
	}
}
