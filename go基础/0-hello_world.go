// package 定义包名 main
package main

//import 引用库 fmt
import "fmt"

// 定义变量(小驼峰命名)
// var 变量名 0-变量类型
var name string = "hello"
var age = 99
var age2 = 66

// func 定义 main 函数
func main() {
	//包名 fmt 调用 Println 方法
	fmt.Println("hello world 终于可以了")
	fmt.Printf("%d\n", age)
	fmt.Println(name)

	age, age2 = age2, age
	fmt.Printf("age=%d,age2=%d\n", age, age2)
	fmt.Printf("age地址：%p\n", &age)
	fmt.Println("age地址的取值", *&age)
}
