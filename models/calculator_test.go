package models

import (
	"fmt"
)

func (ms *ModelSuite) Test_Calculate() {
	tcases := []struct {
		expression string
		result     int
		err        error
	}{
		{"", 0, nil},
		{"1", 1, nil},
		{"2", 2, nil},
		{"20", 20, nil},
		{"1,2", 3, nil},
		{"1,2,3", 6, nil},
		{"1,2,3,5", 11, nil},
		{"2\n3,4", 9, nil},
		{"2,3\n4\n6,10", 25, nil},
		{"//;\n1;2", 3, nil},
		{"//;;\n3;;2", 5, nil},
		{"//:\n5:5", 10, nil},
		{"//:;,\n5:;,9", 14, nil},
		{"//:;,\n5:;,7", 12, nil},
		{"//;:,\n5;:,5\n20", 30, nil},
		{"-1", 0, ErrNumberNotValid},
		{"-1,2,3", 0, ErrNumberNotValid},
	}

	for index, tcase := range tcases {
		calculator := Calculator{
			Expression: tcase.expression,
		}
		err := calculator.Calculate()
		ms.Equal(tcase.err, err, fmt.Sprintf("index: %v", index))
		ms.Equal(tcase.result, calculator.Result, fmt.Sprintf("index: %v", index))
	}
}
