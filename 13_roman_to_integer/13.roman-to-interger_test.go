package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one string
}

type ans struct {
	one int
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
				one: "III",
			},
			a: ans{
				one: 3,
			},
		},
		{
			p: param{
				one: "LVIII",
			},
			a: ans{
				one: 58,
			},
		},
		{
			p: param{
				one: "MCMXCIV",
			},
			a: ans{
				one: 1994,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, romanToInt(p.one), "Input:%v", p)
	}
}
