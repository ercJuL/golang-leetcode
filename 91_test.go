// 一条包含字母A-Z 的消息通过以下映射进行了 编码 ：
//
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
//
// "AAJF" ，将消息分组为 (1 1 10 6)
// "KJF" ，将消息分组为 (11 10 6)
// 注意，消息不能分组为 (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
//
// 给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
//
// 题目数据保证答案肯定是一个 32 位 的整数。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/decode-ways
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"testing"
)

func numDecodings(s string) int {
	n := len(s) + 1
	f := make([]int, n)
	f[0] = 1
	for i := 1; i < n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' <= 26 {
			f[i] += f[i-2]
		}
	}
	return f[n-1]
}

func TestNumDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "基础测试1",
			args: "12",
			want: 2,
		}, {
			name: "基础测试2",
			args: "226",
			want: 3,
		}, {
			name: "基础测试3",
			args: "0",
			want: 0,
		}, {
			name: "错误测试1",
			args: "06",
			want: 0,
		}, {
			name: "错误测试2",
			args: "10",
			want: 1,
		}, {
			name: "错误测试3",
			args: "27",
			want: 1,
		}, {
			name: "错误测试4",
			args: "2101",
			want: 1,
		}, {
			name: "错误测试5",
			args: "1201234",
			want: 3,
		}, {
			name: "错误测试6",
			args: "10011",
			want: 0,
		}, {
			name: "错误测试7",
			args: "123123",
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := numDecodings(tt.args)
				if got != tt.want {
					t.Errorf("numDecodings() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
