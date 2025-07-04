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

func Test58(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: "Hello World",
			},
			a: ans{
				one: 5,
			},
		},
		{
			p: param{
				one: "   fly me   to   the moon  ",
			},
			a: ans{
				one: 4,
			},
		},
		{
			p: param{
				one: "luffy is still joyboy",
			},
			a: ans{
				one: 6,
			},
		},
		{
			p: param{
				one: "a",
			},
			a: ans{
				one: 1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLastWord(p.one), "Input:%v", p)
	}
}