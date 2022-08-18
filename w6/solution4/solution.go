package solution4

import (
	"fmt"
	"time"
)

func Execute() {
	// создаем канал для целых чисел
	ch := make(chan int)

	// запускаем горутину и передаем в нее канал
	go worker(ch)

	// запускаем чтение канала в цикле
	for n := range ch {
		// считаем манулов
		fmt.Printf("%d manul...\n", n)
	}
	// когда канал будет закрыт, цикл завершится
}

func worker(ch chan int) {
	// запускаем цикл и передаем по каналу значение i
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	// закрываем канал
	close(ch)
}
