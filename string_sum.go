package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	errs "github.com/pkg/errors"
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
	if len(input) == 0 {
		return "", fmt.Errorf("could not compute the sum: %w", errorEmptyInput)
	}
	op1, op2, err := splitIntoOperands(input)
	if err != nil {
		//return "", fmt.Errorf("could not compute the sum: %w", err)
		return "", errs.Wrap(err, "could not compute the sum")
	}

	i1, err := strconv.Atoi(op1)
	if err != nil {
		return "", fmt.Errorf("first operand cannot be converted to int: %w", err)
	}

	i2, err := strconv.Atoi(op2)
	if err != nil {
		return "", fmt.Errorf("second operand cannot be converted to int: %w", err)
	}

	return strconv.Itoa(i1 + i2), nil
}

func splitIntoOperands(input string) (op1, op2 string, err error) {
	tokens := strings.Split(input, "+")

	if len(tokens) > 2 {
		return "", "", errorNotTwoOperands
	}

	if len(tokens) == 1 {
		tokensSub := strings.Split(input, "-")

		if len(tokensSub) > 3 || len(tokensSub) <= 1 {
			return "", "", errorNotTwoOperands
		}
		if len(tokensSub) == 3 {
			return "-" + strings.TrimSpace(tokensSub[1]), "-" + strings.TrimSpace(tokensSub[2]), nil
		}
		return strings.TrimSpace(tokensSub[0]), "-" + strings.TrimSpace(tokensSub[1]), nil

	}

	return strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1]), nil
}
