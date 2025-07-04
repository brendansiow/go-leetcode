package goleetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to create a linked list from a slice
func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}

// Helper function to convert linked list to slice for comparison
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

type param struct {
	one []int
	two []int
}

type ans struct {
	one []int
}

type question struct {
	p param
	a ans
}

func Test21(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: param{
				one: []int{1, 2, 4},
				two: []int{1, 3, 4},
			},
			a: ans{
				one: []int{1, 1, 2, 3, 4, 4},
			},
		},
		{
			p: param{
				one: []int{},
				two: []int{},
			},
			a: ans{
				one: []int{},
			},
		},
		{
			p: param{
				one: []int{},
				two: []int{0},
			},
			a: ans{
				one: []int{0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		list1 := createLinkedList(p.one)
		list2 := createLinkedList(p.two)
		result := mergeTwoLists(list1, list2)
		resultSlice := linkedListToSlice(result)
		ast.Equal(a.one, resultSlice, "Input:%v", p)
	}
}