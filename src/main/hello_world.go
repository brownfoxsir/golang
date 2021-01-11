package main // 包，表明代码所在的模块（包）

import (
	"fmt"
	"os"
) // 引入代码依赖

// 功能实现
func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {

		fmt.Println("hello World", os.Args[1])
	}
	os.Exit(0)
}
