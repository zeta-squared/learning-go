package main

func isValidParentheses(s string) bool {
	stack := make([]rune, 0, len(s))
	d := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		default:
			if len(stack) == 0 {
				return false
			}
			open := stack[len(stack)-1]
			if d[open] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
