package solution2

import (
	"fmt"
	"strings"
)

func Execute() {
	s := "snow dog sun" // присваиваем строку переменной

	p := strings.Split(s, " ") // делим строку на слова

	// объявляем новый пустой слайс
	newS := []string{}

	// проходим по слайсу слов в обратном порядке
	for i := len(p) - 1; i >= 0; i-- {
		// добавляем слова в пустой слайс
		newS = append(newS, p[i])
	}

	s = strings.Join(newS, " ") // объединяем слова обратно в строку

	fmt.Println(s)
}
