package solution2

import "fmt"

func Execute() {
	// запрашиваем у пользователя слово
	s := getWord()

	// вызываем функцию переворачивания строки и передаем полученное слово, как аргумент
	r := reverseWord(s)

	fmt.Println(r)
}

func getWord() string {
	fmt.Print("Input: ")
	var s string
	_, err := fmt.Scan(&s) // просим пользователя ввести строку в консоль
	if err != nil {
		panic(err)
	}
	return s
}

func reverseWord(orig string) string {
	// объявляем пустой слайс рун
	reversed := []rune{}

	// конвертируем слово в другой слайс рун
	runes := []rune(orig)

	// с помощью цикла в обратном порядке проходим по слайсу рунами из слова
	for i := len(runes) - 1; i >= 0; i-- {
		// добавляем руны в пустой слайс рун
		reversed = append(reversed, runes[i])
	}

	// преобразовываем слайс рун в строку и возвращаем
	return string(reversed)
}
