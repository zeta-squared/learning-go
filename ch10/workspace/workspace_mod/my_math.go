package workspace_mod

import (
	constraints "golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add is a function to perform addition.
//
// It takes two parameters, both of type int. It returns their sum as type int.
// For more details on addition click [here].
//
// [here]: https://www.mathisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
