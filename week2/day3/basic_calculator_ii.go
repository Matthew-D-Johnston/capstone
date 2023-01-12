package main

/*
Basic Calculator II

Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate results will be in the range of [-2^31, 2^31 - 1].

Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().


Alogrithm:
- create a slice to act as a stack
- Iterate over the chars of the string
	- if a char is an integer then push it onto the stack
	- if its either an addition or subtraction operator, push it onto the stack
	- if it is either a multiplication or division operator then take the last char
		from the stack and perform the respective operation with the next char
		from the string, and push the resulting value onto the stack, then carry
		on from the character that was last used (essentially skip over one char since
		it was used in calculation)

- after the loop, all the multiplication and division operations should be complete
- then just take the stack and pop numbers and operators off, computing the values
  as you go. It will only be addition and subtraction at this point so order of
	operations no longer matters.
*/

import (
	"fmt"
	"strconv"
)

func getNextElement(s string, index int) (string, int) {
	a, r := index, index+1
	fmt.Println(a, r)

	for r < len(s) {
		fmt.Println(a, r)

		if string(s[r]) >= "0" && string(s[r]) <= "9" {
			r++
			continue
		}

		break
	}

	s, index = s[a:r], (r-a)+index
	// fmt.Println(s, index)
	return s, index
}

func calculate(s string) int {
	stack := make([]string, 0)

	for idx := 0; idx < len(s); idx++ {
		char := string(s[idx])
		switch {
		case char == "+" || char == "-":
			stack = append(stack, char)
		case char == " ":
			continue
		case char == "*" || char == "/":
			lastElement, _ := strconv.Atoi(string(stack[len(stack)-1]))
			nextElement, nextIndex := getNextElement(s, idx+1)
			nextElementInt, _ := strconv.Atoi(nextElement)
			// fmt.Println(lastElement, nextElement, nextIndex, nextElementInt)

			var result int
			if char == "*" {
				result = lastElement * nextElementInt
			} else {
				result = lastElement / nextElementInt
			}

			stack = append(stack[0:len(stack)-1], strconv.Itoa(result))
			idx = nextIndex - 1
			continue
		default:
			if len(stack) > 0 {
				lastElement := stack[len(stack)-1]
				if lastElement != "+" && lastElement != "-" {
					lastElement += char
					stack[len(stack)-1] = lastElement
				} else {
					stack = append(stack, char)
				}
			} else {
				stack = append(stack, char)
			}
			// lastElement := stack[len(stack)-1]

			// if lastElement != "+" && lastElement != "-" {
			// 	lastElement += char
			// 	stack[len(stack)-1] = lastElement
			// } else {
			// 	stack = append(stack, char)
			// }
		}

	}
	// fmt.Println(string(stack[0]))
	fmt.Println(stack)
	return 0
}

func main() {
	fmt.Println(calculate("32+10/2-5") == 7)

	// fmt.Println(calculate(" 3/2 ") == 1)
	// fmt.Println(calculate(" 3+5 / 2 ") == 5)
}
