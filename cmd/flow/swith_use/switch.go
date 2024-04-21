package swith_use

func switchInterface(val interface{}) {
	switch val.(type) {
	case int:
		println("val is a int")
	case string:
		println("val is a string")
	default:
		println("Other type")

	}
}

func BasicUse() {
	// 值匹配
	num := 2
	switch num {
	case 1:
		println("One")
	case 2:
		println("Two")
	default:
		println("Other")
	}

	// 类型匹配
	switchInterface(2)
	switchInterface("name")

	// 跳出switch
	newNum := 2
	switch {
	case newNum > 1:
		println("Contains one")
		fallthrough
	case newNum == 2:
		println("Equal two")
	default:
		println("Other")
	}
}
