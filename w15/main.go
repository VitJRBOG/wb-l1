package main

import (
	"fmt"
	"strings"
)

func someFunc() string {
	// генерируем большую строку
	v := createHugeString(1 << 10)

	// преобразуем подстроку в руны, обрезаем и полученный слайс рун преобразуем обратно в строку
	justString := string([]rune(v)[:100])
	return justString
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}

func createHugeString(rowSize int) string {
	return strings.Repeat("g", rowSize)
}
