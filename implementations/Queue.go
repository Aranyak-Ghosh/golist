package implementations

import (
	"encoding/json"
	"sync"

	"github.com/Aranyak-Ghosh/golist/types"
)

type Queue[T any] struct {
	data []T
	mu   sync.Mutex
}

func (q *Queue[T]) UnmarshalJSON(data []byte) error {
	var out []T

	err := json.Unmarshal(data, &out)

	if err != nil {
		return types.UnmarshallError(err)
	} else {
		q.data = out
		return nil
	}
}

func (q *Queue[T]) Length() int {
	return len(q.data)
}

func (q *Queue[T]) Enqueue(entry T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, entry)
}

func (q *Queue[T]) Dequeue() (*T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) == 0 {
		return nil, false
	} else {
		var res = q.data[0]
		q.data = q.data[1:]
		return &res, true
	}
}

func (q *Queue[T]) Peek() (*T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) == 0 {
		return nil, false
	} else {
		var res = q.data[0]
		return &res, true
	}
}
