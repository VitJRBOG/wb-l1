package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10} // объявляем массив с числами

	ch := make(chan int, len(numbers)) // объявляем канал с размером буфера равный количеству чисел в массиве

	wg := sync.WaitGroup{} // объявляем экземпляр WaitGroup

	wg.Add(1)                             // инкрементируем счетчик WaitGroup
	go squareTheNumbers(ch, numbers, &wg) // запускаем в горутине функцию, которая будет считать квадраты чисел из массива

	wg.Add(1)                 // еще раз инкрементируем счетчик WaitGroup
	go sumTheNumbers(ch, &wg) // запускаем в горутине функцию, которая будет получать числа из канала и считать их сумму

	wg.Wait() // ждем, когда обнулится счетчик WaitGroup
}

func squareTheNumbers(ch chan int, numbers [5]int, wg *sync.WaitGroup) {
	internalWg := sync.WaitGroup{} // объявляем локальный экземпляр WaitGroup

	for _, n := range numbers { // перебираем числа из массива
		internalWg.Add(1) // инкрементируем локальный WaitGroup для каждого элемента массива
		go func(n int) {  // запускаем в горутине анонимную функцию, которая будет считать квадрат числа n
			r := n * n                      // вычисляем квадрат n
			fmt.Printf("%d^2 = %d\n", n, r) // выводим результат вычисления квадрата числа n
			ch <- r                         // передаем результат вычислений в канал
			internalWg.Done()               // декрементируем локальный WaitGroup
		}(n) // чтобы избежать захвата переменной n, передаем ее значение в функцию явно
	}

	internalWg.Wait() // ждем обнуления счетчика локального WaitGroup
	close(ch)         // закрываем канал
	wg.Done()         // декрементируем счетчик внешнего WaitGroup
}

func sumTheNumbers(ch <-chan int, wg *sync.WaitGroup) {
	result := 0         // определяем переменную для хранения суммы чисел из канала
	for n := range ch { // в цикле ждем числа из канала
		result += n // прибавляем число из канала
	}

	fmt.Printf("Sum of squares of numbers is: %d\n", result) // выводим результат вычисления суммы чисел из канала
	wg.Done()                                                // декрементируем счетчик WaitGroup
}
