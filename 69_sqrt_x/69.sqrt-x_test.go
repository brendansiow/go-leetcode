package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p param
	a ans
}

func Test69(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: 4,
			},
			a: ans{
				one: 2,
			},
		},
		{
			p: param{
				one: 8,
			},
			a: ans{
				one: 2,
			},
		},
		{
			p: param{
				one: 0,
			},
			a: ans{
				one: 0,
			},
		},
		{
			p: param{
				one: 1,
			},
			a: ans{
				one: 1,
			},
		},
		{
			p: param{
				one: 16,
			},
			a: ans{
				one: 4,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, mySqrt(p.one), "Input:%v", p)
	}
}