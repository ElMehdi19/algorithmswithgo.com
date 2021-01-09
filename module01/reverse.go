package module01

import "strings"
// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//

func pop(stack *[]string) string {
	outputRef := *stack
	lastItem := string(outputRef[len(outputRef) - 1])
	*stack = outputRef[0:len(outputRef) - 1]
	return lastItem
}

func Reverse(word string) string {
	stack := make([]string, len(word))
	for i, char := range word {
		stack[i] = string(char)
	}

	var item string
	reversed := make([]string, len(word))
	for len(stack) != 0 {
		item = pop(&stack)
		reversed = append(reversed, item)
	}

	output := strings.Join(reversed, "")
	return output
}
