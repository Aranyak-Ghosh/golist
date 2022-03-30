package implementations

import (
	"encoding/json"

	"github.com/Aranyak-Ghosh/golist/types"
)

// TypeDef List to array
type List[T any] []T

func (l *List[T]) UnmarshalJSON(data []byte) error {
	var out []T

	err := json.Unmarshal(data, &out)

	if err != nil {
		return types.UnmarshallError(err)
	} else {
		*l = out
		return nil
	}
}

func (l *List[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l)
}

func (l *List[T]) Append(entry T) *List[T] {
	*l = append(*l, entry)
	return l
}

func (l *List[T]) RemoveAt(index int) error {
	Length := len(*l)

	if index >= Length || index < 0 {
		return types.IndexOutOfBoundError()
	} else {
		left := (*l)[:index]
		right := (*l)[index+1:]
		*l = append(left, right...)
		return nil
	}
}

func (l *List[T]) Length() int {
	return len(*l)
}

func (l *List[T]) Map(ex func(T) any) []any {

	var res = make([]any, len(*l))

	for i, v := range *l {
		var tempRes = make(chan any)
		go func() {
			var dat = ex(v)

			tempRes <- dat

		}()
		res[i] = <-tempRes
	}

	return res
}

func (l *List[T]) Filter(ex func(el T) bool) *List[T] {
	var res *List[T] = new(List[T])

	*res = make([]T, l.Length())

	var counter = 0
	for _, v := range *l {
		include := make(chan bool)

		go func(v T) {
			include <- ex(v)
		}(v)
		var inc = <-include
		if inc {
			(*res)[counter] = v
			counter++
		}
	}

	*res = (*res)[:counter]

	return res
}
