package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

// 练习 7.4： strings.NewReader函数通过读取一个string
// 参数返回一个满足io.Reader接口类型的值（和其它值）。
// 实现一个简单版本的NewReader，
// 用它来构造一个接收字符串输入的HTML解析器（§5.2）

func main() {
	_, err := html.Parse(NewReader("<h1Hello</h1>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse err: %v", err)
		os.Exit(1)
	}
}

type StringReader string

func (s *StringReader) Read(p []byte) (int, error) {
	copy(p, *s)
	return len(*s), io.EOF
}

func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}
