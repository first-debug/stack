package stack

type Stack[T any] struct {
	i     int
	slice []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		i:     -1,
		slice: make([]T, 10),
	}
}

func (stack *Stack[T]) Pop() (res T) {
	if stack.i >= 0 {
		res = stack.slice[stack.i]
		stack.i--
	} else {
		panic("Стек пуст!")
	}
	return
}

func (stack *Stack[T]) Push(val T) {
	if stack.i+1 > 9 {
		panic("Стек переполнен!")
	}
	stack.i++
	stack.slice[stack.i] = val
}
