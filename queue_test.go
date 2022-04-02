package golist

import (
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	var q Queue[int]

	var expected = 1
	q.Enqueue(expected)

	var i, _ = q.Peek()
	if *i != expected {
		t.Errorf("Expected %d, got %d", expected, *i)
	}
}

func TestQueuePeek(t *testing.T) {
	var q Queue[int]

	var expected = 1
	q.Enqueue(expected)

	var i, _ = q.Peek()
	if *i != expected {
		t.Errorf("Expected %d, got %d", expected, *i)
	}
}

func TestQueueLenght(t *testing.T) {
	var q Queue[int]

	var expected = 1

	q.Enqueue(5)

	len := q.Length()

	if expected != len {
		t.Errorf("Expected %d, got %d", expected, len)
	}
}

func TestQueueDequeue(t *testing.T) {
	var q Queue[int]

	var expected = 1

	q.Enqueue(expected)

	var i, _ = q.Dequeue()
	if *i != expected {
		t.Errorf("Expected %d, got %d", expected, *i)
	}

	var _, ok = q.Dequeue()

	if ok {
		t.Errorf("Expected false got true")
	}

}
