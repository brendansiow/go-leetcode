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

func Test35(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: []int{1, 3, 5, 6},
				two: 5,
			},
			a: ans{
				one: 2,
			},
		},
		{
			p: param{
				one: []int{1, 3, 5, 6},
				two: 2,
			},
			a: ans{
				one: 1,
			},
		},
		{
			p: param{
				one: []int{1, 3, 5, 6},
				two: 7,
			},
			a: ans{
				one: 4,
			},
		},
		{
			p: param{
				one: []int{1, 3, 5, 6},
				two: 0,
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, searchInsert(p.one, p.two), "Input:%v", p)
	}
}