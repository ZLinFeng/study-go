package main

import (
	"fmt"
	"os"
)

func divide(a int, b int) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return float32(a) / float32(b), nil
}

// 在返回中已经写明了变量的名称
func minAndMax(arr []int) (min int, max int) {
	min, max = arr[0], arr[0]
	for _, item := range arr {
		if item > max {
			max = item
		}
		if item < min {
			min = item
		}
	}
	// 只需要写return即可
	return
}

func getSum() func(num int) int {
	sum := 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

func main() {
	// 函数可以有多个返回值
	if r, e := divide(1, 2); e == nil {
		println(r)
	}

	// 命名返回值
	arr := []int{0, 2, 4, -2, -4}
	min, max := minAndMax(arr)
	fmt.Printf("Min is: %d, Max is %d", min, max)

	// 匿名函数
	func(words string) {
		println(words)
	}("Hello, 这是一个匿名函数, 定义完了就可以直接用")

	// 匿名函数赋值
	greet := func(words string) {
		println(words)
	}
	greet("通过变量调用匿名函数")

	// 闭包: 一个匿名函数和其相关的上下文环境组成的一个整体, 返回这个匿名函数
	f := getSum()
	println("add 1: ", f(1))
	println("add 2: ", f(2))
	println("add 3: ", f(3))
	println("add 4: ", f(4))

	// 函数中的defer关键字, 用于确保调用的函数会在包含它的函数结束时执行
	// 无论是正常结束还是因为错误而退出
	file, err := os.Open("C:\\code\\go\\study-go\\res\\test.txt")
	if err != nil {
		println("Failed to open file.")
	}
	// 结束后会关闭文件
	defer file.Close()
	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		println("Failed to read file.")
	}
	fmt.Println(string(data))

}
