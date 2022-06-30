package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	replacer := strings.NewReplacer(" ", "", "\t", "", "\n", "", "\f", "", "\r", "", "\v", "")
	input = replacer.Replace(input)
	var (
		s []string
	)
	switch len(input) {
	case 0:
		return "", fmt.Errorf("%w", errorEmptyInput)
	case 3:
		s = strings.SplitAfter(input, string(input[0]))
	case 4:
		s = strings.SplitAfter(input, string(input[1]))
	case 5:
		s = strings.SplitAfter(input, string(input[1]))
	case 6:
		s = strings.SplitAfter(input, string(input[2]))
	default:
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}
	s1 := s[0]
	s2 := s[1]
	t1, e1 := strconv.Atoi(s1)
	if e1 != nil {
		return "", fmt.Errorf("%w", e1)
	}
	t2, e2 := strconv.Atoi(s2)
	if e2 != nil {
		return "", fmt.Errorf("%w", e2)
	}
	return strconv.Itoa(t1 + t2), nil
}
