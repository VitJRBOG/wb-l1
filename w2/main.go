package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10} // объявляем массив numbers

	wg := sync.WaitGroup{} // объявляем экземпляр WaitGroup, чтобы основная горутина не завершилась раньше времени

	wg.Add(len(numbers)) // вносим длину numbers в счетчик WaitGroup

	for _, n := range numbers {
		// запускаем в горутине анонимную функцию, вычисляющую квадрат числа n и выводящую его в консоль
		// чтобы не произошло захвата переменной n, передаем ее значение в функцию явно
		go func(n int) {
			fmt.Printf("%d^2 = %d\n", n, n*n)
			wg.Done() // декрементируем счетчик WaitGroup
		}(n)
	}

	wg.Wait() // ждем, пока счетчик WaitGroup не обнулится
}
