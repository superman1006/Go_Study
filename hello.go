package main

import (
	//111111111111111111111111
	"fmt"
	"math/rand"
	//可用于批量包导入
)

/*
	import "fmt"

	import (
	"fmt"
	"math/rand"
	)
*/
// 如果导入包时写."fmt",则但是用fmt包内函数时不需要写 fmt.函数名,可以直接使用函数名调用

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")

	var a = 10
	b := 20 //:起到一个声明变量的作用
	b = 30  //后续对b进行修改的时候不需要加:,因为已经声明过了
	fmt.Println(a)
	fmt.Println(b)
	// = 是赋值， := 是声明变量并赋值

	fmt.Println("My favorite number is", rand.Intn(10)) //输出随机数
}
