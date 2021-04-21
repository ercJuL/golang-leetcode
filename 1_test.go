// 给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 的那两个整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 你可以按任意顺序返回答案。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "基础测试",
			args: args{[]int{2, 7, 11, 15}, 9},
			want: []int{0, 1},
		}, {
			name: "基础测试1",
			args: args{[]int{3, 2, 4}, 6},
			want: []int{1, 2},
		}, {
			name: "基础测试2",
			args: args{[]int{3, 3}, 6},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := twoSum(tt.args.nums, tt.args.target)
				for _, num := range tt.want {
					if !(got[0] == num || got[1] == num) {
						t.Errorf("removeElement() = %v, want %v", got, tt.want)
						break
					}
				}
			},
		)
	}
}
