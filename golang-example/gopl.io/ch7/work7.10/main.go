package main

import (
	"fmt"
	"sort"
)

// 练习 7.10： sort.Interface类型也可以适用在其它地方。
// 编写一个IsPalindrome(s sort.Interface)
// bool函数表明序列s是否是回文序列，
// 换句话说反向排序不会改变这个序列。
// 假设如果!s.Less(i, j) && !s.Less(j, i)
// 则索引i和j上的元素相等。

func main() {
	// 一个字符窜取半 前一般的字符倒序后是等于另一半的
	s := str("abccbb")
	fmt.Println(IsPalindrome(s))
}

type str []byte

func (s str) Len() int           { return len(s) }
func (s str) Less(i, j int) bool { return s[i] < s[j] }
func (s str) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
