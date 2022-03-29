package golist

type baseList[T any] interface {
	Lenght() int
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
	Dequeue() (*T, error)
}

type IStack[T any] interface {
	baseList[T]
	Push(T)
	Pop() (*T, error)
} 