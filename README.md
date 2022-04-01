
# Go List

A library for implementations of basic data structures in Golang. 


## Overview

Implementations include queue, stack and list. All implementations are thread-safe and use mutexes to handle a single process at a time.

A queue provides the following methods in its interface:
- Enqueue: adds an element into the queue in FIFO structure (type-safe).
- Dequeue: removes the first element that was enqueued in a FIFO structure (type-safe)
- Peek: returns the element at the front of the queue, returning a pointer to the element and a boolean value to indicate if the queue is empty.
- Length: returns the length of the queue as an integer.
- UnmarshallJSON: overrides default unmarshal behaviour to allow unmarshalling array to queue
- MarshalJSON: overrides default marshall behaviour to marshall into json array


A list provides the following methods in its interface:
- Append: adds an element into the list in order of entry (type-safe).
- RemoveAt: removes an element at a specified index, takes integer as an argument, returns data type of error.
- Map: creates a map out of the list. 
- Filter: filters from a list using a bool function that satisfies a particular condition.
- UnmarshallJSON: overrides default unmarshal behaviour to allow unmarshalling array to list
- MarshalJSON: overrides default marshall behaviour to marshall into json array


A stack provides the following methods in its interface:
- Push: adds an element into the stack in LIFO structure (type-safe).
- Pop: removes an element from top of the stack in LIFO structure (type-safe).
- Peek: returns the element at the top of the stack in LIFO structure, returning a pointer to the element and a boolean value to indicate if stack is empty.
- Length: returns the length of the stack as an integer.
- UnmarshallJSON: overrides default unmarshal behaviour to allow unmarshalling array to stack
- MarshalJSON: overrides default marshall behaviour to marshall into json array


An error type provides certain error messages for indexes going out of bound, parsing and for internal errors. 

## Example Code

Import this library using the github URL in your import statements

- Stacks
````
var stack implementations.Stack[string]
stack.Push("unicorn")
stack.Push("donkey")
stack.Push("bird")
stack.Push("horse")
stack.Pop() //LIFO
val2, ok2 := stack.Peek()
if !ok2 {
    fmt.Println("Stack is empty")
} else {
    fmt.Println("Value at top of stack is", *val2)
}
fmt.Println("Length of stack is", stack.Length())
````

- Lists 
````
var items implementations.List[int]
items.Append(8)
items.Append(33953)
items.Append(35)
x := items[0]
fmt.Println("Item at index 0 is ", x)
var evens = items.Filter(func(val int) bool {
    return val%2 == 0
})
fmt.Printf("Values that are even: %v\n", *evens)
err := items.RemoveAt(3)
if err != nil {
	fmt.Println("Error: tried to remove at index out of bounds. Operation failed.")
}
fmt.Println("Length of list is", items.Length())
````

- Queues
````
var queue implementations.Queue[int]
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)
queue.Dequeue() //FIFO
val, ok := queue.Peek()
if !ok {
    fmt.Println("Queue is empty")
} else {
	fmt.Println("Value at front of queue is", *val)
}
fmt.Println("Length of queue is", queue.Length())
````



##  Authors

- [@Aranyak-Ghosh](https://github.com/Aranyak-Ghosh)

