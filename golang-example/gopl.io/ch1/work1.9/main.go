package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		netStat := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("请求状态码: %s", netStat)
		fmt.Printf("%s", b)
	}
}
