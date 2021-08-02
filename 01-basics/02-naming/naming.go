package naming

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func (d Dog) Bark(muzzled bool) {
	if muzzled {
		fmt.Println("woof")
	} else {
		fmt.Println("WOOF!!")
	}
}

func Speak(lang string) {
	switch lang {
	case "es":
		fmt.Println("Hola")
	default:
		fmt.Println("Hello")
	}
}

func Color(name string) string {
	switch name {
	case "blue":
		return "#0000ff"
	case "green":
		return "#00ff00"
	case "red":
		return "#ff0000"
	default:
		return "#000000"
	}
}
