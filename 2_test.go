// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0开头。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-two-numbers
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumberSubNode(l1, l2, 0)
}

func addTwoNumberSubNode(l1 *ListNode, l2 *ListNode, n int) *ListNode {
	if l1 == nil && l2 == nil && n == 0 {
		return nil
	}
	l1Val, l2Val := 0, 0
	var l1Next *ListNode = nil
	var l2Next *ListNode = nil
	if l1 != nil {
		l1Val, l1Next = l1.Val, l1.Next
	}
	if l2 != nil {
		l2Val, l2Next = l2.Val, l2.Next
	}
	sum := l1Val + l2Val + n
	v := sum % 10
	return &ListNode{
		Val:  v,
		Next: addTwoNumberSubNode(l1Next, l2Next, sum/10),
	}

}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		args [2][]int
		want []int
	}{
		{
			name: "基础测试",
			args: [2][]int{{2, 4, 3}, {5, 6, 4}},
			want: []int{7, 0, 8},
		}, {
			name: "基础测试1",
			args: [2][]int{{0}, {0}},
			want: []int{0},
		}, {
			name: "基础测试2",
			args: [2][]int{{9, 9, 9, 9, 9, 9, 9}, {9, 9, 9, 9}},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := addTwoNumbers(sliceToNond(tt.args[0]), sliceToNond(tt.args[1]))
				for _, num := range tt.want {
					if !(got != nil && got.Val == num) {
						t.Errorf("removeElement() = %v, want %v", got, tt.want)
						break
					}
					got = got.Next
				}
			},
		)
	}
}

func sliceToNond(s []int) *ListNode {
	if len(s) <= 0 {
		return nil
	}
	return &ListNode{Val: s[0], Next: sliceToNond(s[1:])}
}
