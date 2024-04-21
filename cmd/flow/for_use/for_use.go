package for_use

import "fmt"

func BasicUse() {
	// 传统的for循环
	for i := 0; i < 5; i++ {
		println("传统循环: ", i)
	}

	// 类型while的for循环
	i := 0
	for i < 5 {
		println("for like while: ", i)
		i++
	}

	// 无限循环
	for {
		println("无限循环")
		break
	}
}

func AdvancedUse() {
	// 遍历数组、字符串、map
	lastName := "LastName"
	for i, char := range lastName {
		fmt.Printf("Char %d is: %c\n", i, char)
	}

	// 除了continue和break外, 还可以使用标签
outerloop:

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue outerloop
			}
			fmt.Printf("i is: %d, j is: %d\n", i, j)
		}
	}

	// 多变量初始化
	for i, j := 0, 100; i < 2; i, j = i+1, j-50 {
		fmt.Printf("i is: %d, j is %d\n", i, j)
	}
}
