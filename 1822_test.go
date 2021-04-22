// 已知函数signFunc(x) 将会根据 x 的正负返回特定值：
//
// 如果 x 是正数，返回 1 。
// 如果 x 是负数，返回 -1 。
// 如果 x 是等于 0 ，返回 0 。
// 给你一个整数数组 nums 。令 product 为数组 nums 中所有元素值的乘积。
//
// 返回 signFunc(product) 。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sign-of-the-product-of-an-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"testing"
)

func arraySign(nums []int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		switch {
		case nums[i] < 0:
			r++
		case nums[i] == 0:
			return 0
		}
	}
	if r%2 == 0 {
		return 1
	} else {
		return -1
	}
}

func TestArraySign(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "基础测试",
			args: []int{-1, -2, -3, -4, 3, 2, 1},
			want: 1,
		}, {
			name: "基础测试1",
			args: []int{1, 5, 0, 2, -3},
			want: 0,
		}, {
			name: "基础测试2",
			args: []int{-1, 1, -1, 1, -1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := arraySign(tt.args)
				if got != tt.want {
					t.Errorf("arraySign() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
