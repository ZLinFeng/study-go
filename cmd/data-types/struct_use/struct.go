package struct_use

type Person struct {
	Name string
	Age  int
}

func (p Person) greeting() string {
	return "Hello World, I'm " + p.Name
}

type Employee struct {
	// 将Person 嵌入到 Employee
	Person
	Salary int8
}

func AdvancedUse() {
	// 结构体方法
	edison := Person{Name: "Edison", Age: 30}
	println(edison.greeting())

	// 结构体的嵌入和模拟继承
	e := Employee{
		Person: Person{
			Name: "Edison", Age: 30,
		},
		Salary: 100,
	}
	println(e.greeting(), "的薪水是: ", e.Salary)

}
