package interfaces

type baseList[T any] interface {
	Length() int
	Append(T)
}

type IList[T any] interface {
	baseList[T]
	Map(func(T) any) []any
	Filter(func(T) bool) IList[T]
}

type IQueue[T any] interface {
	baseList[T]
	Enqueue(T)
	Dequeue() (*T, bool)
	Peek() (*T, bool)
}

type IStack[T any] interface {
	baseList[T]
	Push(T)
	Pop() (*T, bool)
	Peek() (*T, bool)
}
