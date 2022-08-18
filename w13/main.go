package main

import "fmt"

func main() {
	array := [9]int{4, 2, 1, 4, 6, 1, 2, 8, 6} // объявляем массив с целыми числами

	fmt.Println(array)

	array[1], array[5] = array[5], array[1] // присваиваем элементу 1 значение из 5, а элементу 5 значение из 1

	fmt.Println(array)
}
