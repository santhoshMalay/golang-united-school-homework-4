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
	errEmpty := fmt.Errorf("%s", errorEmptyInput)
	errMoreLess := fmt.Errorf("%s", errorNotTwoOperands)

	input = strings.ReplaceAll(input, " ", "")
	strLen := len(input)
	if strLen == 0 {
		return "", errEmpty
	}

	addSymbol := strings.ContainsAny(input, "-")
	subSymbol := strings.ContainsAny(input, "+")

	if addSymbol == false && subSymbol == false {
		return "", errMoreLess
	}

	startIndex := 0
	if strings.HasPrefix(input, "-") == true {
		startIndex = 1
	}

	num1, digitsInNum1, sum, operand := 0, 0, 0, 0
	for x := startIndex; x < strLen; x++ {
		if input[x] >= 48 && input[x] <= 57 {
			digitsInNum1 = x + 1
		} else {
			operand = int(input[x])
			num2, err1 := strconv.Atoi(string(input[x+1 : strLen]))
			if err1 != nil {
				return "", errMoreLess
			}
			sum += num2
			break
		}
	}
	num1, err1 := strconv.Atoi(string(input[0:digitsInNum1]))
	if err1 != nil {
		return "", errMoreLess
	}
	if operand == 43 {
		sum += num1
	} else if operand == 45 {
		sum = num1 - sum
	} else {
		return "", errMoreLess
	}

	return strconv.Itoa(sum), nil
}
