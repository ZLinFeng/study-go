package pointer

import "fmt"

func modify(p *int) {
	*p = 21
}

func BasicType() {
	// 地址
	a := 10
	p := &a
	println("a is: ", a)
	println("a address is: ", &a)
	fmt.Println("p value is: ", p)

	// 通过指针访问和修改
	var b int = 11
	var bp *int = &b
	println("直接访问b: ", b)
	println("通过指针访问b: ", *bp)

	*bp = 12
	println("通过指针修改后的b: ", *bp)

	// 通过函数来直接修改
	c := 10
	modify(&c)
	println("通过函数修改之后: ", c)

	// 指针的默认值
	var defaultP *int
	fmt.Println(defaultP)
}
