package solution2

import (
	"fmt"
	"time"
)

func Execute() {
	n := getSecondsToEnd() // получаем от пользователя количество секунд, через которое программа завершится

	ch := make(chan int)              // объявляем канал для передачи данных
	chOfEnd := make(chan interface{}) // объявляем канал для отправки сигнала о завершении работы

	go inputData(ch, chOfEnd) // запускаем в горутине функцию передачи данных в канал
	go outputData(ch)         // запускаем в горутине функцию чтения данных из канала

	time.Sleep(time.Duration(n) * time.Second) // запускаем обратный отсчет до завершения работы программы

	var x struct{} // объявляем экземпляр пустой структуры
	chOfEnd <- x   // передаем экземпляр пустой структуры в канал для завершения работы
	close(chOfEnd) // закрываем канал для отправки сигнала о завершении работы
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

// inputData отправляет данные через канал ch и принимает сигнал о завершении работы из канала chOfEnd
func inputData(ch chan int, chOfEnd <-chan interface{}) {
InputLoop: // указываем метку, чтобы цикл можно было завершить из блока select
	for i := 0; ; i++ {
		select {
		case <-chOfEnd: // проверяем, не пришло ли что-то из канала для завершения работы
			close(ch)       // если пришло, то закрываем канал на отправку данных
			break InputLoop // и завершаем цикл
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
