package arrays

import (
	"errors"
)

type Array struct {
	items []interface{}
	size  int
}

func New() *Array {
	newArray := Array{items: make([]interface{}, 10)}
	return &newArray
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) Push(item interface{}) {
	a.items[a.size] = item
	a.size++
}

func (a *Array) Pop() interface{} {
	a.size--
	return a.items[a.size]
}

func (a Array) Filter(filterFunction func(item interface{}, index int, items []interface{}) bool) *Array {
	newArray := New()
	for i := 0; i < a.Size(); i++ {
		if filterFunction(a.items[i], i, a.items) {
			newArray.Push(a.items[i])
		}
	}
	return newArray
}

func (a Array) Map(mapFunction func(item interface{}, index int, items []interface{}) bool) *Array {
	newArray := New()
	for i := 0; i < a.Size(); i++ {
		newArray.Push(mapFunction(a.items[i], i, a.items))
	}
	return newArray
}

func (a Array) Every(everyFunction func(item interface{}, index int, items []interface{}) bool) bool {

	for i := 0; i < a.Size(); i++ {
		if !everyFunction(a.items[0], i, a.items) {
			return false
		}
	}
	return true
}

func (a Array) Some(someFunction func(item interface{}, index int, items []interface{}) bool) bool {

	for i := 0; i < a.Size(); i++ {
		if someFunction(a.items[0], i, a.items) {
			return true
		}
	}
	return false
}

func (a Array) ForEach(function func(item interface{}, index int, items []interface{}) bool) {
	for i := 0; i < a.Size(); i++ {
		function(a.items[i], i, a.items)
	}
}

func (a *Array) At(position int) (interface{}, error) {
	if position < 0 || position >= a.Size() {
		return nil, errors.New("Position is not within the range of items in the array")
	}

	return a.items[position], nil
}

func (a Array) Concat(a2 Array) *Array {
	newArray := New()
	var count int = 0
	for ; count < a.Size(); count++ {
		newArray.Push(a.items[count])
	}
	count = 0
	for ; count < a2.Size(); count++ {
		newArray.Push(a.items[count])
	}
	return newArray
}

func (a Array) Find(item interface{}) (int, error) {
	for i := 0; i < a.Size(); i++ {
		if a.items[i] == item {
			return i, nil
		}
	}
	return -1, errors.New("Item does not exist in array")
}
