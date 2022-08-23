package solution1

import "fmt"

func Execute() {
	fmt.Print("Input: ")
	var s string
	_, err := fmt.Scan(&s) // просим пользователя ввести строку в консоль
	if err != nil {
		panic(err)
	}

	r := []rune(s) // переводим строку в руны, чтобы можно было обращаться к символам по индексу

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // меняем символы местами
	}

	fmt.Println("Output: " + string(r))
}
