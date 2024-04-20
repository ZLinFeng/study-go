package interface_use

type Animal interface {
	Speak() string
}

type Dog struct {
	Word string
}

type Cat struct {
	LittleWord string
}

func (d Dog) Speak() string {
	return d.Word + ", Wang~"
}

func (c Cat) Speak() string {
	return c.LittleWord + ", Miao~"
}

func speak(a Animal) {
	println(a.Speak())
}

func BasicUse() {
	dog := Dog{Word: "汪"}
	cat := Cat{LittleWord: "喵"}

	speak(dog)
	speak(cat)
}
