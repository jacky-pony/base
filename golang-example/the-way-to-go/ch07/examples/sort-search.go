package main

import (
    "fmt"
    "sort"
)

// sort包的排序、搜索示例
func main() {
    // 搜索
    a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 7  
    
    i := sort.Search(len(a), func(i int) bool { return a[i] >= x})
    if i < len(a) && a[i] == x {
        fmt.Printf("found %d at index %d in %v\n", x, i, a)
    } else {
        fmt.Printf("%d not foud in %v\n", x, a)
    }
}
