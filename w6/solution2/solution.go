package solution2

import (
	"fmt"
	"time"
)

func Execute() {
	// создаем канал для строк
	ch := make(chan string)
	// и канал для пустых структур (им будем завершать горутину)
	endCh := make(chan struct{})

	// запускаем горутины
	go worker1(ch, endCh)
	go worker2(endCh)

	// в цикле читаем канал для строк и выводим их в консоль
	for s := range ch {
		fmt.Println(s)
	}
}

func worker1(ch chan string, endCh chan struct{}) {
	fmt.Println("Worker1: I was born!")
	for {
		select {
		// проверяем, не пришло ли что-то из канала завершения
		case <-endCh:
			// если пришло, то выводим последнее сообщение
			ch <- "Worker1: Oh, no..."
			// и закрываем канал для строк
			close(ch)
			return
		default:
			// если нет, то пишем в канал для строк дальше
			ch <- "Worker1: I'm still alive!"
		}
		time.Sleep(1 * time.Second)
	}
}

func worker2(endCh chan struct{}) {
	time.Sleep(5 * time.Second)
	fmt.Println("Worker2: I will destroy you!")
	// отправляем в канал для завершения пустую структуру
	endCh <- struct{}{}
	// и закрываем канал для завершения
	close(endCh)
}
