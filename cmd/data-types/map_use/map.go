package map_use

func BasicUse() {

	// 通过make创建map
	firstMap := make(map[string]string)

	firstMap["first_name"] = "Edison"
	firstMap["last_name"] = "Zhu"

	// 检查是否存在

	if lastName, ok := firstMap["last_name"]; ok {
		println("Last name is ", lastName)
	} else {
		println("....")
	}

	// 遍历map
	for key, value := range firstMap {
		println("Key is: ", key, " and Value is: ", value)
	}

	// 删除元素
	delete(firstMap, "last_name")
	if lastName, ok := firstMap["last_name"]; ok {
		println("Last name is ", lastName)
	} else {
		println("....")
	}

}

type Person struct {
	Age  int
	Name string
}

func AdvancedUse() {
	name2Person := make(map[string]Person)

	name2Person["bob"] = Person{Age: 23, Name: "Bob"}
	name2Person["edison"] = Person{Age: 24, Name: "Edison"}

	anotherBob := name2Person["bob"]
	anotherBob.Age = 25             // 这样修改并没有改变map中的值
	name2Person["bob"] = anotherBob // 通过重新赋值来修改
	if per, ok := name2Person["bob"]; ok {
		println(per.Name, "'s new age is: ", per.Age)
	}
}
