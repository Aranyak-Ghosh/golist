package golist

import (
	"github.com/Aranyak-Ghosh/golist/implementation"
	"github.com/Aranyak-Ghosh/golist/interfaces"
)

type IQueue[T any] interfaces.IQueue[T]
type IList[T any] interfaces.IList[T]
type IStack[T any] interfaces.IStack[T]

type Queue[T any] implementation.Queue[T]
type List[T any] implementation.List[T]
type Stack[T any] implementation.Stack[T]
