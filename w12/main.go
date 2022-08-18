package main

import "fmt"

func main() {
	set := [5]string{"cat", "cat", "dog", "cat", "tree"} // объявляем массив со строками

	items := map[string]struct{}{} // объявляем мапу пустых структур

	for _, value := range set { // проходим по массиву
		if _, ok := items[value]; ok { // если элемент массива присутствует в мапе в качестве ключа, то завершаем итерацию
			continue
		}

		items[value] = struct{}{} // иначе добавляем элемент массива в мапу в качестве ключа и передаем ей пустую структуру
	}

	fmt.Println(items)
}
