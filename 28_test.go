// 实现strStr()函数。
//
// 给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
//
//
//
// 说明：
//
// 当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//
// 对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/implement-strstr
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import (
	"testing"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	needleByte := []byte(needle)
	needleByteLen := len(needleByte)
	haystackByte := []byte(haystack)
	haystackLen := len(haystack)

	needleIndex := 0    // 目标索引
	matchIndex := 0     // 匹配索引
	lastMatchIndex := 0 // 最后匹配索引
	matched := false    // 是否匹配
	for matchIndex < haystackLen {
		if haystackLen-matchIndex < needleByteLen-needleIndex {
			return -1
		}
		if needleIndex >= needleByteLen {
			matched = true
			break
		}
		if haystackByte[matchIndex] == needleByte[needleIndex] {
			if needleIndex == 0 {
				lastMatchIndex = matchIndex
			}
			needleIndex++
			matchIndex++
			matched = true
		} else {
			matched = false
			if needleIndex > 0 {
				matchIndex = lastMatchIndex + 1
			} else {
				matchIndex++
			}
			needleIndex = 0
		}
	}
	if !matched {
		return -1
	}
	return lastMatchIndex
}

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "基础测试",
			args: args{"hello", "ll"},
			want: 2,
		}, {
			name: "基础测试",
			args: args{"aaaaa", "bba"},
			want: -1,
		}, {
			name: "基础测试",
			args: args{"", ""},
			want: 0,
		}, {
			name: "错误测试1",
			args: args{"aaa", "aa"},
			want: 0,
		}, {
			name: "错误测试2",
			args: args{"aaa", "aaaa"},
			want: -1,
		}, {
			name: "错误测试3",
			args: args{"mississippi", "issip"},
			want: 4,
		}, {
			name: "错误测试4",
			args: args{"aabaaabaaac", "aabaaac"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got := strStr(tt.args.haystack, tt.args.needle)
				if got != tt.want {
					t.Errorf("strStr() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
