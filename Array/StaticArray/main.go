package main

import (
	"errors"
	"fmt"
	"unsafe"
)

var (
	ErrInvalidIndex    = errors.New("invalid index")
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrArrayFull       = errors.New("array full")
)

type Array[T any] struct {
	delta *T
	data  []T
	size  int
}

func New[T any](size int) Array[T] {
	data := make([]T, size)

	return Array[T]{
		delta: &data[0],
		data:  data,
		size:  size,
	}
}

func (arr *Array[T]) GetValueAsPointer(index int) T {
	size := unsafe.Sizeof(*arr.delta)

	address := unsafe.Add(unsafe.Pointer(arr.delta), uintptr(index)*size)

	return *(*T)(address)
}

func main() {
	arr := New[int16](5)

	arr.data[3] = 6544       // set data
	fmt.Println(arr.data[2]) // get data

	fmt.Println(arr.GetValueAsPointer(3))

}
