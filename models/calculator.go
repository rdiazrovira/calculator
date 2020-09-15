package models

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrNumberNotValid is returned when a string does not represent a positive number
	ErrNumberNotValid = errors.New("number not valid")
)

type Calculator struct {
	Expression string `json:"expression"`
	Result     int    `json:"result"`
}

func (c *Calculator) Calculate() error {
	expression := c.Expression
	if expression == "" {
		return nil
	}

	separator := ","
	if strings.HasPrefix(expression, "//") && strings.Contains(expression, "\n") {
		separator = expression[2:strings.Index(expression, "\n")]
		expression = expression[strings.Index(expression, "\n")+1:]
	}

	expression = strings.ReplaceAll(expression, "\n", separator)

	if strings.Contains(expression, separator) {
		var sum int
		for _, numberStr := range strings.Split(expression, separator) {
			num, err := strconv.Atoi(numberStr)
			if err != nil || num < 0 {
				return ErrNumberNotValid
			}
			sum += num
		}
		c.Result = sum
		return nil
	}

	num, err := strconv.Atoi(expression)
	if err != nil || num < 0 {
		return ErrNumberNotValid
	}
	c.Result = num
	return nil
}
