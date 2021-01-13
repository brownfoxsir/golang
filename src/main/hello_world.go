package main // 包，表明代码所在的模块（包）

import (
	"fmt"
	"os"
) // 引入代码依赖

// 功能实现
func main() {
	//fmt.Println(os.Args)
	if len(os.Args) > 1 {

		fmt.Println("hello World", os.Args[1])
	} else {
		fmt.Print("1. fmt.Print：我加了换行符\n")
		fmt.Println("2. fmt.Println：我会自动换行")
		fmt.Print("3. fmt.Print：我不会换行 ")
		fmt.Println("4. fmt.Println：我和3在同一行")
		os.Exit(0)
	}
}
