package main

import (
	"errors"
	"fmt"
)

type Array[T any] struct {
	data     []T
	capacity int
	size     int
}

func New[T any](capacity uint) *Array[T] {
	return &Array[T]{
		data:     make([]T, capacity),
		capacity: int(capacity),
		size:     0,
	}
}

func (a *Array[T]) Get(index int) (T, error) {
	if index < 0 || index >= a.size {
		errors.New("index out of range")
	}
	return a.data[index], nil
}

func (a *Array[T]) Put(index int, value T) error {
	if index < 0 {
		return errors.New("index out of range")
	}

	if index >= a.capacity {
		newArray := make([]T, index*2)
		copy(newArray, a.data)

		a.data = newArray
		a.data[index] = value
		return nil
	}

	a.data[index] = value
	return nil
}

func main() {
	arr := New[int](5)
	fmt.Println(arr.data)
	fmt.Println(arr.Put(2, 2))
	fmt.Println(arr.Get(2))
	fmt.Println(arr.data)
	fmt.Println(arr.Put(30, 3))
	fmt.Println(arr.data)
}
