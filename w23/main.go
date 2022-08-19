package main

import "fmt"

func main() {
	// объявляем слайс с числами
	s := []int{3, 2, 1, 3, 2, 2, 1, 4, 5, 6}

	// вызываем функцию удаления элемента по индексу
	s = removeItemByIndex(s, 4)

	fmt.Println(s)
}

func removeItemByIndex(s []int, i int) []int {
	// объявляем пустой слайс
	m := []int{}
	// присваиваем пустому слайсу значения старого слайса до элемента i
	m = append(m, s[:i]...)
	// и значения после элемента i
	m = append(m, s[i+1:]...)
	return m
}
