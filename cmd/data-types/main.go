package main
import "fmt"

func main()  {

	/* 
	go语言中的数据类型有 基本类型、复合类型、引用类型和接口类型
	 */
	 fmt.Println("首先是基本类型: ")
	 var isValid bool = false
	 fmt.Println("bool类型的申明: var isValid bool = ", isValid)

	 // 字符型
	 var sampleRune rune = '你' // rune 实际上是int32
	 var sampleByte byte = 'A' //  byte 实际上是uint8
	 fmt.Println("Sample Rune: (Unicode character): ", sampleRune)
	 fmt.Println("Sample byte: (ASCII character): ", sampleByte)
	 
	 // 字符串
	 var helloWord string = "Hello World."
	 fmt.Println(helloWord)
}