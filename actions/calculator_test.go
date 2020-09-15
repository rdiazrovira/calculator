package actions

import (
	"fmt"
	"net/http"
	"net/url"
)

func (as *ActionSuite) Test_CalculateShow() {
	res := as.HTML("/calculate").Get()
	as.Equal(http.StatusOK, res.Code)

	body := res.Body.String()
	as.Contains(body, `<h2>StringCalculator for</h2>`)
	as.Contains(body, `<input class=" form-control" name="Expression" placeholder="Type your expression..." type="text" value="" />`)
	as.Contains(body, `<button type="submit" class="btn btn-primary">CALCULATE</button>`)
}

func (as *ActionSuite) Test_Calculate() {
	tcases := []struct {
		expression string
		result     int
	}{
		{"", 0},
		{"1", 1},
		{"2", 2},
		{"20", 20},
		{"1,2", 3},
		{"1,2,3", 6},
		{"1,2,3,5", 11},
		{"2\n3,4", 9},
		{"2,3\n4\n6,10", 25},
		{"//;\n1;2", 3},
		{"//;;\n3;;2", 5},
		{"//:\n5:5", 10},
		{"//:;,\n5:;,9", 14},
		{"//:;,\n5:;,7", 12},
		{"//;:,\n5;:,5\n20", 30},
		{"-1", 0},
		{"-1,2,3", 0},
	}

	for index, tcase := range tcases {
		res := as.HTML("/calculate").Post(url.Values{
			"Expression": []string{tcase.expression},
		})
		as.Equal(http.StatusOK, res.Code)

		body := res.Body.String()
		as.Contains(body, fmt.Sprintf("<span class=\"text-center font-100\">%v</span>", tcase.result), fmt.Sprintf("index: %v", index))
	}
}
