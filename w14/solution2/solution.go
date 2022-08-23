package solution2

import "fmt"

func Execute() {
	// объявляем несколько переменных разных типов
	var i int
	var s string
	var b bool
	ch := make(chan struct{})

	// объявляем слайс для данных произвольного типа и передаем в него значения разных типов
	array := []interface{}{i, s, b, ch}

	// перебираем этот слайс
	for _, x := range array {
		// и передаем каждый его элемент в функцию проверки типа
		checkType(x)
	}
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int ")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan struct{}:
		fmt.Println("chan struct{}")
	default:
		fmt.Println("unhandled type")
	}
}
