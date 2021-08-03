package datas

import (
	"fmt"
	"testing"
)

func testSlicesAreEqual(got, expected []interface{}) error {
	if len(got) != len(expected) {
		return fmt.Errorf("len(got) != len(expected)")
	}
	for i := range got {
		if got[i] != expected[i] {
			return fmt.Errorf("got[%d] = %d, expected %d", i, got[i], expected[i])
		}
	}
	return nil
}

func TestSlice(t *testing.T) {
	l := SynglyLinkedList{}
	l.PushFront(10)
	l.PushFront(20)
	l.PushBack(30)
	l.PushBack(40) // l = [20, 10, 30, 40]

	got := l.Slice()
	expected := []interface{}{20, 10, 30, 40}

	err := testSlicesAreEqual(got, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestIsEmpty(t *testing.T) {
	l := SynglyLinkedList{}

	if !l.IsEmpty() || l.Size() != 0 {
		t.Error("List is not empty.")
	}
}

func TestPushFront(t *testing.T) {
	l := SynglyLinkedList{}
	l.PushFront(10)
	l.PushFront(20)
	l.PushFront(30)

	got := l.Slice()
	expected := []interface{}{30, 20, 10}

	err := testSlicesAreEqual(got, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPopFront(t *testing.T) {
	l := SynglyLinkedList{}
	l.PushFront(10)
	l.PushFront(20)
	l.PushFront(30)

	got1 := l.PopFront()
	if got1 != 30 {
		t.Errorf("PopFront() = %d, expected %d", got1, 30)
	}

	got2 := l.Slice()
	expected := []interface{}{20, 10}

	err := testSlicesAreEqual(got2, expected)
	if err != nil {
		t.Errorf(err.Error())
	}

	l.PopFront()
	l.PopFront()
	got3 := l.Slice()
	expected = []interface{}{}
	err = testSlicesAreEqual(got3, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPushBack(t *testing.T) {
	l := SynglyLinkedList{}
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	got := l.Slice()
	expected := []interface{}{10, 20, 30}

	err := testSlicesAreEqual(got, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPopBack(t *testing.T) {
	l := SynglyLinkedList{}
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	got1 := l.PopBack()
	if got1 != 30 {
		t.Errorf("PopBack() = %d, expected %d", got1, 30)
	}

	got2 := l.Slice()
	expected := []interface{}{10, 20}

	err := testSlicesAreEqual(got2, expected)
	if err != nil {
		t.Errorf(err.Error())
	}

	l.PopBack()
	l.PopBack()
	got3 := l.Slice()
	expected = []interface{}{}
	err = testSlicesAreEqual(got3, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSplice(t *testing.T) {
	l := SynglyLinkedList{}
	l.PushFront(10)
	l.PushFront(20)
	l.PushBack(30)
	l.PushBack(40) // l = [20, 10, 30, 40]

	l.Splice(1, 2, 100, 200, 300) // l = [20, 100, 200, 300, 40]

	got := l.Slice()
	expected := []interface{}{20, 100, 200, 300, 40}

	err := testSlicesAreEqual(got, expected)
	if err != nil {
		t.Errorf(err.Error())
	}

	l.Splice(0, l.Size())
	got = l.Slice()
	expected = []interface{}{}
	err = testSlicesAreEqual(got, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGet(t *testing.T) {
	l := SynglyLinkedList{}
	l.PushFront(10)
	l.PushFront(20)
	l.PushBack(30)
	l.PushBack(40) // l = [20, 10, 30, 40]

	got := l.Get(0)
	expected := 20
	if got != expected {
		t.Errorf("l[0] = %d, expected %d", got, expected)
	}

	got = l.Get(1)
	expected = 10
	if got != expected {
		t.Errorf("l[1] = %d, expected %d", got, expected)
	}

	got = l.Get(-2)
	expected = 30
	if got != expected {
		t.Errorf("l[-2] = %d, expected %d", got, expected)
	}

	got = l.Get(-1)
	expected = 40
	if got != expected {
		t.Errorf("l[-1] = %d, expected %d", got, expected)
	}
}
