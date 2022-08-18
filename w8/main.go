package main

import "fmt"

func main() {
	var n int64 = 100000
	var i int64 = 3

	r := shift(n, i)

	fmt.Println(r)
}

func shift(n, i int64) int64 {
	// создаем маску
	m := int64(1 << i)

	// выполняем дизъюнкцию
	r := n | m

	return r
}
