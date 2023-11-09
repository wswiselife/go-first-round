package main

import "fmt"

func foo(num1 int, str1 string) (int, string) {
	num := num1 + 10
	str := str1 + "hello"
	return num, str
}

func main() {
	bae, bat := foo(10, "yan")
	fmt.Println("bae = ", bae, "bat = ", bat)
}

