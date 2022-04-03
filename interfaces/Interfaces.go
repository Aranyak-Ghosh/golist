package interfaces

type baseList[T any] interface {
	Length() int
}

type IList[T any] interface {
	baseList[T]
	Append(T) IList[T]
	Map(func(T) any) IList[any]
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
