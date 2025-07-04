package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	one string
	two string
}

type ans struct {
	one string
}

type question struct {
	p param
	a ans
}

func Test67(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: "11",
				two: "1",
			},
			a: ans{
				one: "100",
			},
		},
		{
			p: param{
				one: "1010",
				two: "1011",
			},
			a: ans{
				one: "10101",
			},
		},
		{
			p: param{
				one: "0",
				two: "0",
			},
			a: ans{
				one: "0",
			},
		},
		{
			p: param{
				one: "1",
				two: "111",
			},
			a: ans{
				one: "1000",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addBinary(p.one, p.two), "Input:%v", p)
	}
}