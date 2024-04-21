package main

import (
	"fmt"
	"os"
)

func openfile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed open file: %w", err)
	}
	defer file.Close()
	return file, nil
}

func main() {
	// 错误处理是: 以返回值的方式使开发者显式的处理
	file, err := os.Open("./file.txt")
	if err != nil {
		println(fmt.Errorf("failed open file: %w", err))
	}
	defer file.Close()

	// 自定义异常 可以直接 error.News("")
	if _, err := openfile("./file.txt"); err != nil {
		_, ok := err.(*os.PathError)
		if ok {
			println("file not found")
		} else {
			// 因为自定义的, 所以是一个error wrap
			println(err.Error())
		}
	}

	// 通过defer和recover来捕获异常
	a, b := 10, 0
	defer func() {
		err := recover()
		if err != nil {
			println("出现异常了")
		}
	}()
	println(a / b)
}
