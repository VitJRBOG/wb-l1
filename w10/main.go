package main

import "fmt"

func main() {
	// объявляем слайс с дробями
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// объявляем пустую мапу для дробей
	m := map[int][]float32{}

	// 5 раз прогоняем слайс с дробями через функцию вычисления шага
	for i := -5; i <= 5; i++ {
		calculateStep(temperatures, m, i)
	}

	fmt.Println(m)
}

func calculateStep(temperatures []float32, m map[int][]float32, step int) {
	// объявляем пустой слайс с дробями
	var numbers []float32
	// в цикле перебираем внешний слайс с дробями
	for _, t := range temperatures {
		// вычисляем количество десятков в целой части дроби
		l := t / 10
		// и проверяем, равна ли десятая часть шагу
		if int(l) == step {
			// если равна, то добавляем число в локальный слайс с дробями
			numbers = append(numbers, t)
		}
	}
	// если в локальном слайсе с дробями что-нибудь есть
	if len(numbers) > 0 {
		// то добавляем это в мапу, в качестве ключа указывая шаг умноженный на 10
		m[step*10] = numbers
	}
}
