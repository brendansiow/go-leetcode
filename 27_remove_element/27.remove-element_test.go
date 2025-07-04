package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one []int
	two int
}

type ans struct {
	one int
}

type question struct {
	p param
	a ans
}

func Test27(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: []int{3, 2, 2, 3},
				two: 3,
			},
			a: ans{
				one: 2,
			},
		},
		{
			p: param{
				one: []int{0, 1, 2, 2, 3, 0, 4, 2},
				two: 2,
			},
			a: ans{
				one: 5,
			},
		},
		{
			p: param{
				one: []int{1},
				two: 1,
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		inputCopy := make([]int, len(p.one))
		copy(inputCopy, p.one)
		result := removeElement(inputCopy, p.two)
		ast.Equal(a.one, result, "Input:%v", p)
		// Check that the first 'result' elements don't contain the target value
		if result > 0 {
			for i := 0; i < result; i++ {
				ast.NotEqual(p.two, inputCopy[i], "Input:%v, found target value at index %d", p, i)
			}
		}
	}
}