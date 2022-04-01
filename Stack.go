package golist

import (
	"encoding/json"
	"sync"

	"github.com/Aranyak-Ghosh/golist/types"
)

type Stack[T any] struct {
	data []T
	mu   sync.Mutex
}

func (s *Stack[T]) UnmarshalJSON(data []byte) error {
	var out []T

	err := json.Unmarshal(data, &out)

	if err != nil {
		return types.UnmarshallError(err)
	} else {
		s.data = out
		return nil
	}
}

func (s *Stack[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.data)
}

func (s *Stack[T]) Push(entry T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, entry)
}

func (s *Stack[T]) Pop() (*T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.data) == 0 {
		return nil, false
	} else {
		var res = s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return &res, true
	}
}

func (s *Stack[T]) Length() int {
	return len(s.data)
}

func (s *Stack[T]) Peek() (*T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.data) == 0 {
		return nil, false
	} else {
		var res = s.data[len(s.data)-1]
		return &res, true
	}
}
