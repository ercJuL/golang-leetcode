// 给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-element
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"testing"
)

func removeElement(nums []int, val int) int {
	slowP := 0 // 慢指针
	fastP := 0 // 快指针
	for ; fastP < len(nums); fastP++ {
		if nums[fastP] == val {
			nums[fastP] = 0
			continue
		}
		nums[slowP] = nums[fastP]
		slowP++
	}
	return slowP
}

func TestRemoveelement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "基础测试",
			args: args{[]int{3, 2, 2, 3}, 3},
			want: args{[]int{2, 2}, 2},
		},
		{
			name: "错误测试1",
			args: args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			want: args{[]int{0, 1, 3, 0, 4}, 5},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := removeElement(tt.args.nums, tt.args.val)
				if got != tt.want.val {
					t.Errorf("removeElement() = %v, want %v", got, tt.want)
				}
				for i, num := range tt.want.nums {
					if tt.args.nums[i] != num {
						t.Errorf("removeElement() = %v, want %v", got, tt.want)
						break
					}
				}
			},
		)
	}
}
