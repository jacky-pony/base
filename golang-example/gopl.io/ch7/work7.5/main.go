package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 练习 7.5： io包里面的LimitReader函数接收
// 一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节
// 但是当读完n个字节后就表示读到文件结束的Reader。
// 实现这个LimitReader函数：

func main() {
	lr := LimitReader(strings.NewReader("abcd"), 2)
	b, err := ioutil.ReadAll(lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
	}
	fmt.Printf("%s\n", b)
}

type LimitedReader struct {
	underlyingReader io.Reader
	remainBytes      int64
}

func (r *LimitedReader) Read(p []byte) (n int, err error) {
	if r.remainBytes <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.remainBytes {
		p = p[:r.remainBytes]
	}
	n, err = r.underlyingReader.Read(p)
	r.remainBytes -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
