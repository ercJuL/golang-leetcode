// 传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。
//
// 传送带上的第 i个包裹的重量为weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
//
// 返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"testing"
)

func shipWithinDays(weights []int, D int) int {
	totalWeight, maxWeight := 0, 0
	for i := 0; i < len(weights); i++ {
		totalWeight += weights[i]
		if maxWeight < weights[i] {
			maxWeight = weights[i]
		}
	}
	needMinWeight := totalWeight / D
	if maxWeight > needMinWeight {
		needMinWeight = maxWeight
	}
	needDay := 1
	for {
		wt := 0
		for i := 0; i < len(weights); i++ {
			wt += weights[i]
			if wt > needMinWeight {
				wt = weights[i]
				needDay++
			}

		}
		if needDay <= D {
			return needMinWeight
		}
		needDay = 1
		needMinWeight++
	}
}

func TestShipWithinDays(t *testing.T) {
	type args struct {
		weights []int
		D       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "基础测试",
			args: args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
			want: 15,
		}, {
			name: "基础测试2",
			args: args{[]int{3, 2, 2, 4, 1, 4}, 3},
			want: 6,
		}, {
			name: "基础测试3",
			args: args{[]int{1, 2, 3, 1, 1}, 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := shipWithinDays(tt.args.weights, tt.args.D)
				if got != tt.want {
					t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
