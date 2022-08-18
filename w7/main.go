package main

import (
	"fmt"
	"sync"
)

func main() {
	// объявляем пустую мапу
	m := map[int]string{}

	// создаем WaitGroup
	wg := new(sync.WaitGroup)
	// и Mutex
	mutex := new(sync.Mutex)

	// увеличиваем счетчик WaitGroup на 3
	wg.Add(3)
	// запускаем горутины и передаем в них мапу и ссылки на WaitGroup и Mutex
	go firstWriter(m, wg, mutex)
	go secondWriter(m, wg, mutex)
	go thirdWriter(m, wg, mutex)

	// ждем обнуления счетчика WaitGroup
	wg.Wait()
	fmt.Println(m)
}

func firstWriter(m map[int]string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	// объявляем слайс с 3 строками
	symbs := []string{"a", "b", "c"}

	// запускаем цикл, в котором будем записывать содержимое слайса в мапу
	for i := 0; i < 3; i++ {
		// блокируем Mutex, чтобы остальные горутины не могли в него писать
		mutex.Lock()
		// добавляем в мапу запись: ключ - число i, значение - i-й элемент в слайсе
		m[i] = symbs[i]
		// разблокируем Mutex
		mutex.Unlock()
	}
	// уменьшаем счетчик WaitGroup
	wg.Done()
}

func secondWriter(m map[int]string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	// объявляем слайс с 3 строками
	symbs := []string{"d", "e", "f"}

	for i := 3; i < 6; i++ {
		// блокируем Mutex, чтобы остальные горутины не могли его заблокировать и, как следствие, писать в него
		mutex.Lock()
		// добавляем в мапу запись: ключ - число i, значение - i-й элемент в слайсе
		m[i] = symbs[i-3]
		// разблокируем Mutex
		mutex.Unlock()
	}
	// уменьшаем счетчик WaitGroup
	wg.Done()
}

func thirdWriter(m map[int]string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	// объявляем слайс с 3 строками
	symbs := []string{"g", "h", "i"}

	for i := 6; i < 9; i++ {
		// блокируем Mutex, чтобы остальные горутины не могли его заблокировать и, как следствие, писать в него
		mutex.Lock()
		// добавляем в мапу запись: ключ - число i, значение - i-й элемент в слайсе
		m[i] = symbs[i-6]
		// разблокируем Mutex
		mutex.Unlock()
	}
	// уменьшаем счетчик WaitGroup
	wg.Done()
}
