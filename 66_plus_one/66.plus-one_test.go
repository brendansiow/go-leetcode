package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one []int
}

type ans struct {
	one []int
}

type question struct {
	p param
	a ans
}

func Test66(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: []int{1, 2, 3},
			},
			a: ans{
				one: []int{1, 2, 4},
			},
		},
		{
			p: param{
				one: []int{4, 3, 2, 1},
			},
			a: ans{
				one: []int{4, 3, 2, 2},
			},
		},
		{
			p: param{
				one: []int{9},
			},
			a: ans{
				one: []int{1, 0},
			},
		},
		{
			p: param{
				one: []int{9, 9},
			},
			a: ans{
				one: []int{1, 0, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, plusOne(p.one), "Input:%v", p)
	}
}