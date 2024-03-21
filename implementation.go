package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

const (
	operators string = "+-*/^"
)

func CalculatePrefix(input string) (int, error) {
	var stack Stack
	var num int
	var err error

	input = strings.Trim(input, " ")

	for i := len(input) - 1; i >= 0; i-- {
		clearSpaces(input, &i)

		if unicode.IsDigit(rune(input[i])) {
			num, i, err = getNumber(input, i)
			if err != nil {
				return -1, err
			}
			stack.Push(num)
		} else {
			if !strings.Contains(operators, string(input[i])) {
				return -1, fmt.Errorf("Error. Problem with opertator")
			}
			if len(stack) < 2 {
				return -1, fmt.Errorf("Error. Missing arguments or many operators")
			}
			o1, _ := stack.Pop()
			o2, _ := stack.Pop()
			switch input[i] {
			case '+':
				stack.Push(o1 + o2)
			case '-':
				stack.Push(o1 - o2)
			case '*':
				stack.Push(o1 * o2)
			case '/':
				if o2 == 0 {
					return -1, fmt.Errorf("Error. Divide by zero")
				}
				stack.Push(o1 / o2)
			case '^':
				stack.Push(int(math.Pow(float64(o1), float64(o2))))
			}
		}
	}
	if stack.Length() != 1 {
		return -1, fmt.Errorf("Error. Missing arguments or many operators")
	}
	res, _ := stack.Pop()
	return res, nil
}

func clearSpaces(data string, i *int) {
	for string(data[*i]) == " " {
		*i--
	}
}

func getNumber(data string, i int) (int, int, error) {
	j := i
	for unicode.IsDigit(rune(data[j])) {
		j--
		if j == -1 {
			break
		}
	}
	j++
	result, err := strconv.Atoi(data[j : i+1])
	if err != nil {
		return -1, -1, fmt.Errorf("Error. NaN")
	}
	return result, j, err
}

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Length() int {
	return len(*s)
}

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
