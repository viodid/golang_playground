package main

import (
	"errors"
	"fmt"
	"unsafe"
)

// == EX01 ==
type IntOrFloat interface {
	uint8 | uint16 | uint32 | uint64 |
		int8 | int16 | int32 | int64 |
		float32 | float64
}

func double[T IntOrFloat](n T) T {
	return n * 2
}

// == EX02 ==
type Printable interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
	fmt.Stringer
}

func printValue[P Printable](p P) {
	fmt.Println(p)
}

type myFloat float32

func (mf myFloat) String() string {
	return fmt.Sprintf("%f", mf)
}

// == EX03 ==
type List[T comparable] struct {
	value T
	next  *List[T]
}

func newList[T comparable](t T) *List[T] {
	return &List[T]{
		value: t,
		next:  nil,
	}
}

// In theory, l would be a deep copy of the obj (I need to clarify this point, as in C there is no deep copy).
// So no node would be appended in the caller obj.
func (l *List[T]) Add(t T) {
	fmt.Printf("%p - %p\n", l.next, l)
	for l.next != nil {
		l = l.next
	}
	l.next = newList(t)
	fmt.Printf("size struct: %d\n", unsafe.Sizeof(l))
	fmt.Println("inside:", l.next)
}

func (l *List[T]) Insert(t T, i int) error {
	n := newList(t)
	if i == 0 {
		n.next = l
		return nil
	}
	c := 0
	for l != nil && c != i-1 {
		l = l.next
		c++
	}
	if l == nil {
		return errors.New("index out of limits")
	}
	nextNode := l.next
	l.next = n
	n.next = nextNode
	return nil
}

func (l List[T]) Index(t T) int {
	return -1
}

func main() {
	var x float32 = 2.3
	var y int8 = 42
	fmt.Println(double(x), double(y))
	var z myFloat = 42.42
	printValue(z)

	l := &List[string]{}
	l.Add("hey")
	for l != nil {
		fmt.Println(l)
		l = l.next
	}
}
