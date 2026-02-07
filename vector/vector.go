package vector

import "fmt"

type Vector[T any] struct {
	arr      []T
	capacity int
}

func NewVector[T any]() *Vector[T] {

	initialCap := 1

	return &Vector[T]{
		arr:      make([]T, 0, initialCap),
		capacity: initialCap,
	}

}

func (v *Vector[T]) Size() int {
	return len(v.arr)
}

func (v *Vector[T]) Capacity() int {
	return v.capacity
}

func (v *Vector[T]) IsEmpty() bool {
	return v.Size() == 0
}

func (v *Vector[T]) At(index int) T {
	if index < 0 || index >= v.Size() {
		panic("Index out of range")
	}
	return v.arr[index]
}

func (v *Vector[T]) Push(elem T) {
	// check if we need to resize
	if v.Size()+1 > v.Capacity() {
		// double the capacity
		v.capacity *= 2
		// create resized array
		newArr := make([]T, v.Size(), v.capacity)
		// copy old to new
		copy(newArr, v.arr)
		v.arr = newArr
	}
	v.arr = append(v.arr, elem)
}

func (v *Vector[T]) Insert(index int, elem T) error {
	// check bounds
	if index < 0 || index > v.Size() {
		return fmt.Errorf("Index %d out of bounds.", index)
	}

	// add new space at the end
	var zero T
	v.Push(zero)

	// shift elements to the right, starting from the end
	for i := v.Size() - 1; i > index; i-- {
		v.arr[i] = v.arr[i-1]
	}

	// insert the element
	v.arr[index] = elem

	return nil
}

func (v *Vector[T]) Prepend(element T) {
	v.Insert(0, element)
}

func (v *Vector[T]) Pop() (T, error) {
	var zero T
	if v.Size() == 0 {
		return zero, fmt.Errorf("Can't pop from an empty vector.")
	}
	lastElem := v.arr[v.Size()-1]
	v.arr[v.Size()-1] = zero
	v.arr = v.arr[:v.Size()-1]
	return lastElem, nil
}

func (v *Vector[T]) Delete(index int) error {
	if index < 0 || index >= v.Size() {
		return fmt.Errorf("index out of range")
	}

	// shift elements to the left
	for i := index; i < v.Size()-1; i++ {
		v.arr[i] = v.arr[i+1]
	}

	// zero out the last element and shrink
	var zero T
	v.arr[v.Size()-1] = zero
	v.arr = v.arr[:v.Size()-1]

	return nil
}

// Shallow copy - caller can modify item.Data
func (v *Vector[T]) GetItems() []T {
	result := make([]T, len(v.arr))
	copy(result, v.arr)
	return result
}
