package main

import (
	"fmt"
	"vector/vector"
)

func ShowVectorInfo[T any](v *vector.Vector[T]) {
	fmt.Printf("Size: %d, cap: %d\n", v.Size(), v.Capacity())
	for _, e := range v.GetItems() {
		fmt.Print(e, ", ")
	}
	fmt.Println()
}

func main() {
	s := vector.NewVector[int]()
	fmt.Println("[Push] Adding 1")
	s.Push(1)
	ShowVectorInfo(s)
	fmt.Println("[Push] Adding 2")
	s.Push(2)
	fmt.Println("[Push] Adding 3")
	s.Push(3)
	ShowVectorInfo(s)
	fmt.Println("[Insert] Inserting 4 at 2")
	s.Insert(2, 4)
	ShowVectorInfo(s)
	fmt.Println("[Insert] Inserting 5 at 4")
	s.Insert(4, 5)
	ShowVectorInfo(s)
	fmt.Println("[Insert] Inserting 6 at 5")
	s.Insert(5, 6)
	fmt.Println("[Prepend] Prepending 7")
	s.Prepend(7)
	fmt.Println("[Prepend] Prepending 8")
	s.Prepend(8)
	ShowVectorInfo(s)
	fmt.Println("[Prepend] Prepending 9")
	s.Prepend(9)
	ShowVectorInfo(s)
	fmt.Println("[Pop] Removing last element")
	s.Pop()
	ShowVectorInfo(s)
	fmt.Println("[Delete] Deleting at 2")
	s.Delete(2)
	ShowVectorInfo(s)
	fmt.Println("[Delete] Deleting at 4")
	s.Delete(4)
	ShowVectorInfo(s)
}
