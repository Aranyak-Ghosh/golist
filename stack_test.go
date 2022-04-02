package golist

import (
	"testing"
)

func TestStackEnqueue(t *testing.T) {
	var q Stack[int]

	var expected = 1
	q.Push(expected)

	var i, _ = q.Peek()
	if *i != expected {
		t.Errorf("Expected %d, got %d", expected, *i)
	}
}

func TestStackPeek(t *testing.T) {
	var q Stack[int]

	var expected = 1
	q.Push(expected)

	var i, _ = q.Peek()
	if *i != expected {
		t.Errorf("Expected %d, got %d", expected, *i)
	}
}

func TestStackLenght(t *testing.T) {
	var q Stack[int]

	var expected = 1

	q.Push(5)

	len := q.Length()

	if expected != len {
		t.Errorf("Expected %d, got %d", expected, len)
	}
}

func TestStackDequeue(t *testing.T) {
	var q Stack[int]

	var expected = 1

	q.Push(expected)

	var i, _ = q.Pop()
	if *i != expected {
		t.Errorf("Expected %d, got %d", expected, *i)
	}

	var _, ok = q.Pop()

	if ok {
		t.Errorf("Expected false got true")
	}

}
