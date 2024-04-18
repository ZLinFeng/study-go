package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func basicType() {
	// go语言中的数据类型有 基本类型、复合类型、引用类型和接口类型
	fmt.Println("首先是基本类型: ")
	var isValid bool = false
	fmt.Println("bool类型的申明: var isValid bool = ", isValid)

	// 字符型
	var sampleRune rune = '你' // rune 实际上是int32
	var sampleByte byte = 'A' //  byte 实际上是uint8
	fmt.Printf("Sample Rune: (Unicode character): 实际类型是 %T, 占 %d个字节\n", sampleRune, unsafe.Sizeof(sampleRune))
	fmt.Println("Sample byte: (ASCII character): ", sampleByte)

	// 字符串
	var helloWord string = "Hello World."
	fmt.Println(helloWord)

	// 进制及转化
	fmt.Println("二进制以0b开头: ", 0b101010)
	fmt.Println("八进制以0开头: ", 052)
	fmt.Println("十六进制以0x开头: ", 0x2A)

	binary := strconv.FormatInt(int64(42), 2)
	println("用二进制表示: ", binary)
	octal := strconv.FormatInt(int64(42), 8)
	println("用八进制表示: ", octal)
	hexadecimal := strconv.FormatInt(int64(42), 16)
	println("用十六进制表示: ", hexadecimal)

	binaryStr := "101010"
	octalStr := "52"
	hexadecimalStr := "2a"
	parseBinary, _ := strconv.ParseInt(binaryStr, 2, 64)
	fmt.Printf("二进制 %s 表示的数字是: %d\n", binaryStr, parseBinary)
	parseOctal, _ := strconv.ParseInt(octalStr, 8, 64)
	fmt.Printf("八进制 %s 表示的数字是: %d\n", octalStr, parseOctal)
	parseHexadecimal, _ := strconv.ParseInt(hexadecimalStr, 16, 64)
	fmt.Printf("十六进制 %s 表示的数字是: %d\n", hexadecimalStr, parseHexadecimal)
}

func compositeType() {
	// 数组
	var firstArray [5]int = [5]int{0, 1, 2, 3, 4}
	println("数组的第3个元素: ", firstArray[2])
	println("遍历所有元素: ")
	for i := 0; i < len(firstArray); i++ {
		fmt.Printf("第 %d 个元素是 %d\n", i, firstArray[i])
	}
}

func main() {
	basicType()
	compositeType()
}
