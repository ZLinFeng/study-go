package array

import "fmt"

func BasicUse() {
	// 数组
	var firstArray [5]int = [5]int{0, 1, 2, 3, 4}
	println("数组的第3个元素: ", firstArray[2])
	println("遍历所有元素: ")
	for i := 0; i < len(firstArray); i++ {
		fmt.Printf("第 %d 个元素是 %d\n", i, firstArray[i])
	}
	var secondArr [3][6]int = [3][6]int{
		{0, 1, 2, 3, 4, 5},
		{6, 7, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
	}
	for rowIndex, row := range secondArr {
		for colIndex, col := range row {
			fmt.Printf("第%d行的第%d个元素是: %d\n", rowIndex+1, colIndex+1, col)
		}
	}
}

func printArr(arr [3]int) {
	for _, i := range arr {
		println(i)
	}
}

func modifyArr(arr *[3]int) {
	arr[2] = 5
}

func AdvancedUse() {
	// 数组作为函数的参数
	arr := [3]int{3, 4, 5}
	printArr(arr)

	// 大数组可以传递数组的指针
	modifyArr(&arr)
	println("修改后")
	printArr(arr)

	// 与切片的互操作
	println("切片")
	a := [6]int{1, 2, 3, 4, 5, 6}
	b := a[1:3]
	for _, i := range b {
		println(i)
	}
}
