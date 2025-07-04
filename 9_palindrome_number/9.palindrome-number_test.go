package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one int
}

type ans struct {
	one bool
}

type question struct {
	p param
	a ans
}

func Test9(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: 121,
			},
			a: ans{
				one: true,
			},
		},
		{
			p: param{
				one: -121,
			},
			a: ans{
				one: false,
			},
		},
		{
			p: param{
				one: 10,
			},
			a: ans{
				one: false,
			},
		},
		{
			p: param{
				one: 0,
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isPalindrome(p.one), "Input:%v", p)
	}
}