package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one string
}

type ans struct {
	one bool
}

type question struct {
	p param
	a ans
}

func Test20(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: "()",
			},
			a: ans{
				one: true,
			},
		},
		{
			p: param{
				one: "()[]{}",
			},
			a: ans{
				one: true,
			},
		},
		{
			p: param{
				one: "(]",
			},
			a: ans{
				one: false,
			},
		},
		{
			p: param{
				one: "([)]",
			},
			a: ans{
				one: false,
			},
		},
		{
			p: param{
				one: "{[]}",
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isValid(p.one), "Input:%v", p)
	}
}