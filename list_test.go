package golist

import (
	"fmt"
	"testing"
)

func TestListEnqueue(t *testing.T) {
	var q List[int]

	var expected = 1
	q.Append(expected)

	var i = q[0]
	if i != expected {
		t.Errorf("Expected %d, got %d", expected, i)
	}
}

func TestListLenght(t *testing.T) {
	var q List[int]

	var expected = 1

	q.Append(5)

	len := q.Length()

	if expected != len {
		t.Errorf("Expected %d, got %d", expected, len)
	}
}

func TestListRemoveAt(t *testing.T) {
	var q List[int]

	var expected = 1

	q.Append(expected)

	var i = q[0]
	if i != expected {
		t.Errorf("Expected %d, got %d", expected, i)
	}

	q.RemoveAt(0)

	if q.Length() != 0 {
		t.Errorf("Expected %d, got %d", 0, q.Length())
	}

}

func TestListFilter(t *testing.T) {
	var q *List[int] = new(List[int])

	*q = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var expected = []int{2, 4, 6, 8, 10}

	q = (q.Filter(func(x int) bool {
		return x%2 == 0
	})).(*List[int])

	for i := range expected {
		if (*q)[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], (*q)[i])
			break
		}
	}
}

func TestListMap(t *testing.T) {
	var q *List[int] = new(List[int])

	*q = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var expected = []string{"plsWork-1", "plsWork-2", "plsWork-3", "plsWork-4", "plsWork-5", "plsWork-6", "plsWork-7", "plsWork-8", "plsWork-9", "plsWork-10"}

	res := (q.Map(func(x int) any {
		return fmt.Sprintf("plsWork-%d", x)
	})).(*List[any])

	for i := range expected {
		if (*res)[i] != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], (*res)[i])
			break
		}
	}
}
