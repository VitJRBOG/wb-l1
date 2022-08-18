package main

import (
	"fmt"
	"sync"
)

func main() {
	// создаем канал для чисел
	numCh := make(chan int)
	// и канал для результата умножения
	multiplicationCh := make(chan int)

	// создаем WaitGroup
	wg := new(sync.WaitGroup)

	// увеличиваем счетчик WaitGroup на 3
	wg.Add(3)
	// и запускаем горутины
	go inputNumbers(wg, numCh)
	go multiplyTheNumbers(wg, numCh, multiplicationCh)
	go printNumbers(wg, multiplicationCh)

	// ждем обнуления счетчика WaitGroup
	wg.Wait()
}

func inputNumbers(wg *sync.WaitGroup, ch chan int) {
	// создаем слайс с числами
	numbers := []int{9, 2, 51, 1, 10, 28}

	// запускаем цикл и перебираем слайс с числами
	for _, n := range numbers {
		// передаем числа в канал для чисел
		ch <- n
	}
	// после завершения перебора слайса закрываем канал
	close(ch)
	// и уменьшаем счетчик WaitGroup
	wg.Done()
}

func multiplyTheNumbers(wg *sync.WaitGroup, numCh <-chan int, squareCh chan int) {
	// в цикле читаем канал с числами
	for n := range numCh {
		// полученное число умножаем на 2
		s := n * 2
		// и передаем в канал для результата умножения
		squareCh <- s
	}
	// после закрытия канала с числами закрываем канал для результатов умножения
	close(squareCh)
	// и уменьшаем счетчик WaitGroup
	wg.Done()
}

func printNumbers(wg *sync.WaitGroup, ch <-chan int) {
	// в цикле читаем канал с результатами умножения
	for n := range ch {
		// полученные числа выводим в консоль
		fmt.Println(n)
	}
	// после закрытия канала уменьшаем счетчик WaitGroup
	wg.Done()
}
