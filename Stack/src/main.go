package main

import "github.com/ManoloTonto1/Semester4-week3/stack"

func main() {
	s := stack.Stack[int]{}
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Show()
	s.Pop()
	s.Show()
	s.Pop()
	s.Show()
}
