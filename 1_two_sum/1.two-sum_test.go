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
	one []int
}

type question struct {
	p param
	a ans
}

func Test1(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: []int{2, 7, 11, 15},
				two: 9,
			},
			a: ans{
				one: []int{0, 1},
			},
		},
		{
			p: param{
				one: []int{3, 2, 4},
				two: 6,
			},
			a: ans{
				one: []int{1, 2},
			},
		},
		{
			p: param{
				one: []int{3, 3},
				two: 6,
			},
			a: ans{
				one: []int{0, 1},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, twoSum(p.one, p.two), "Input:%v", p)
	}
}
