package main

import "fmt"

func main() {
	// int  --> %d
	// float --> %f
	// String --> %s
	var a = 123.2
	var b = "2023-03-01"
	var url1 = "Code = %f & endDate = %s"
	var target_url = fmt.Sprintf(url1, a, b)
	fmt.Println(target_url)

	//fmpl快捷键
	fmt.Println()

	// %d 表示整型数字，%s 表示字符串
	var c = 123
	var d = "2023-03-01"
	var url2 = "Code = %d & endDate = %s"
	fmt.Printf(url2, c, d)

}
