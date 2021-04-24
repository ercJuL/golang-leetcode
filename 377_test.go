// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
// 题目数据保证答案符合 32 位整数范围。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum-iv
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"sort"
	"testing"
)

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	sumMap := map[int]int{}
	return combinationSum4E(sumMap, nums, target)
}

func combinationSum4E(sumMap map[int]int, nums []int, target int) int {
	sum := 0
	oldSum, ok := sumMap[target]
	if ok {
		return oldSum
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		} else if nums[i] == target {
			sum += 1
			break
		} else if nums[i] < target {
			sum += combinationSum4E(sumMap, nums, target-nums[i])
		}
	}
	sumMap[target] = sum
	return sum
}

func TestCombinationSum4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "基础测试",
			args: args{[]int{1, 2, 3}, 4},
			want: 7,
		}, {
			name: "基础测试1",
			args: args{[]int{9}, 3},
			want: 0,
		}, {
			name: "错误测试1",
			args: args{[]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}, 10},
			want: 9,
		}, {
			name: "超时测试1",
			args: args{[]int{2, 1, 3}, 35},
			want: 746448752,
		}, {
			name: "超时测试1",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
				33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65,
				66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98,
				99, 100}, 31},
			want: 746448752,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := combinationSum4(tt.args.nums, tt.args.target)
				if got != tt.want {
					t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func BenchmarkCombinationSum4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		combinationSum4([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
			33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65,
			66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98,
			99, 100}, 31)
	}
}
