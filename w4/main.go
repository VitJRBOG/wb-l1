package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	n := getNumberOfWorkers() // получаем количество воркеров

	dataCh := make(chan interface{}, n)   // объявляем канал для произвольных данных
	breakCh := makeInterruptingListener() // вызываем функцию, объявляющую канал для отслеживания команды прерывания

	for i := n; i > 0; i-- {
		go printData(dataCh) // запускаем n горутин, в которых будут исполняться функции вывода данных из канала в консоль
	}

DataInserting: // указываем метку перед циклом, чтобы его можно было завершить из блока select
	for i := 0; ; i++ {
		select {
		case <-breakCh: // проверяем, не пришло ли что из канала для отслеживания прерывания
			close(dataCh)       // если пришло, то закрываем канал для данных
			close(breakCh)      // и канал для отслеживания прерывания
			break DataInserting // затем выходим из цикла
		default: // если не пришло, то передаем по каналу для данных число
			dataCh <- i
		}
	}
}

// getNumberOfWorkers запрашивает у пользователя целое число - количество воркеров;
// если ввести не int или не ввести ничего, то будет паника
func getNumberOfWorkers() int {
	fmt.Print("Number of workers: ")
	n := 0
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}

	return n
}

// makeInterruptingListener объявляет канал, в который будет передаваться команда на прерывание, и возвращает этот канал
func makeInterruptingListener() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	return ch
}

// printData слушает канал и выводит в консоль данные, которые приходят из канала;
// после закрытия канала функция завершает работу
func printData(ch <-chan interface{}) {
	for d := range ch {
		fmt.Println(d)
	}
}
