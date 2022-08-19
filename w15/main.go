package main

import "strings"

var justString string

func someFunc() {
	// генерируем большую строку
	v := createHugeString(1 << 10)

	// преобразуем подстроку в руны, обрезаем и полученный слайс рун преобразуем обратно в строку
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
}

func createHugeString(rowSize int) string {
	return strings.Repeat("g", rowSize)
}
