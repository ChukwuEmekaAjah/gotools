package arrays

import (
	"fmt"
	"testing"
)

func TestNewArray(t *testing.T) {
	fmt.Println([]byte("abc"))
	newArray := New()
	if newArray.Size() != 0 {
		t.Errorf("New array size should be zero but we got %d", newArray.Size())
	}
}

func TestPushFunc(t *testing.T) {
	newArray := New()
	var items [4]int = [4]int{2, 31, 38, 4}
	for i := 0; i < len(items); i++ {
		newArray.Push(items[i])
	}

	if newArray.Size() != len(items) {
		t.Errorf("New array length should be equal to number of inserted items but got %d", newArray.Size())
	}
}

func TestPopFunc(t *testing.T) {
	newArray := New()
	var items [4]int = [4]int{2, 31, 38, 4}
	for i := 0; i < len(items); i++ {
		newArray.Push(items[i])
	}
	newArray.Pop()
	newArray.Pop()

	if newArray.Size() != (len(items) - 2) {
		t.Errorf("New array length should be equal to number of remaining items after pop but got %d", newArray.Size())
	}
}

func TestAtFunc(t *testing.T) {
	newArray := New()
	var items [4]int = [4]int{2, 31, 38, 4}
	for i := 0; i < len(items); i++ {
		newArray.Push(items[i])
	}

	position := 2
	fourthItem, err := newArray.At(position)

	if err != nil {
		t.Error("Array 'At' function returns error for known position")
	}

	if fourthItem != items[position] {
		t.Errorf("Expected item at position %d to be %d but got %d", position, items[position], fourthItem)
	}
}

func TestEveryFunc(t *testing.T) {
	newArray := New()
	var items [4]int = [4]int{2, 31, 38, 4}
	for i := 0; i < len(items); i++ {
		newArray.Push(items[i])
	}

	response := newArray.Every(func(item interface{}, index int, items []interface{}) bool {
		if item.(int) > 0 {
			return true
		}
		return false
	})

	if !response {
		t.Error("Expected all items in the array to be greater than 0")
	}
}

func TestSomeFunc(t *testing.T) {
	newArray := New()
	var items [4]int = [4]int{-2, 31, 38, 4}
	for i := 0; i < len(items); i++ {
		newArray.Push(items[i])
	}

	response := newArray.Some(func(item interface{}, index int, items []interface{}) bool {
		if item.(int) < 0 {
			return true
		}
		return false
	})

	if !response {
		t.Error("Expected some items in the array to be less than 0")
	}
}

func TestFindFunc(t *testing.T) {
	newArray := New()
	var items [4]int = [4]int{-2, 31, 38, 4}
	for i := 0; i < len(items); i++ {
		newArray.Push(items[i])
	}

	item, _ := newArray.Find(-2)

	if item == -1 {
		t.Error("Expected find to retrieve item in array")
	}
}
