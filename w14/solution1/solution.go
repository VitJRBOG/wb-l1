package solution1

import (
	"fmt"
	"reflect"
)

func Execute() {
	// объявляем несколько переменных разных типов
	var i int
	var s string
	var b bool
	ch := make(chan struct{})

	// объявляем слайс для данных произвольного типа и передаем в него значения разных типов
	array := []interface{}{i, s, b, ch}

	// перебираем этот слайс
	for _, item := range array {
		// и передаем каждый его элемент в функцию проверки типа
		checkType(item)
	}

}

func checkType(x interface{}) {
	fmt.Println(reflect.TypeOf(x)) // выводим тип переменной x
}
