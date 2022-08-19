package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun" // присваиваем строку переменной

	p := strings.Split(s, " ") // делим строку на слова

	// проходим по слайсу слов и меняем их местами
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}

	s = strings.Join(p, " ") // объединяем слова обратно в строку

	fmt.Println(s)
}
