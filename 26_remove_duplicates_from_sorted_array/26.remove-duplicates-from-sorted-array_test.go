package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one []int
}

type ans struct {
	one int
	expectedArray []int
}

type question struct {
	p param
	a ans
}

func Test26(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: []int{1, 1, 2},
			},
			a: ans{
				one: 2,
				expectedArray: []int{1, 2},
			},
		},
		{
			p: param{
				one: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			a: ans{
				one: 5,
				expectedArray: []int{0, 1, 2, 3, 4},
			},
		},
		{
			p: param{
				one: []int{1},
			},
			a: ans{
				one: 1,
				expectedArray: []int{1},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		inputCopy := make([]int, len(p.one))
		copy(inputCopy, p.one)
		result := removeDuplicates(inputCopy)
		ast.Equal(a.one, result, "Input:%v", p)
		// Check that the first 'result' elements match expected array
		if result > 0 {
			ast.Equal(a.expectedArray, inputCopy[:result], "Input:%v", p)
		}
	}
}