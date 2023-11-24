package main

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	var stack []rune

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			}
		case ']':
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			}
		case '}':
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			}
		default:
			continue
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func isValidV2(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	stack := []rune{}
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
}
