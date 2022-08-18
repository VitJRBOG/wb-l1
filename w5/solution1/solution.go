package solution1

import (
	"context"
	"fmt"
	"time"
)

func Execute() {
	n := getSecondsToEnd() // получаем от пользователя количество секунд до завершения работы

	ch := make(chan int) // объявляем канал для передачи данных

	// создаем экземпляр контекста с таймаутом в n секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	// запускаем горутину, которая будет передавать данные через канал ch
	go inputData(ctx, ch)

	// вызываем функцию, которая будет читать данные из канала ch
	outputData(ch)
}

// getSecondsToEnd запрашивает у пользователя целое число; если введенное значение будет не int или будет пустым, то случится паника
func getSecondsToEnd() int {
	fmt.Print("Seconds to end: ")
	n := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	return n
}

// inputData отправляет данные через канал ch до тех пор, пока активен контекст ctx
func inputData(ctx context.Context, ch chan int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): // проверяем, не завершился ли контекст
			close(ch) // если пришло, то закрываем канал на отправку данных
			return
		default: // если сигнала о завершении работы нет, то передаем число i в канал для данных
			ch <- i
		}
	}
}

// outputData принимает из канала данные; если канал закроется, то функция завершит работу
func outputData(ch <-chan int) {
	for data := range ch { // ждем от канала для данных данные
		fmt.Println(data) // и выводим их
	}
}
