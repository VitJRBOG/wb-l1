package solution3

import (
	"fmt"
	"sync"
	"time"
)

func Execute() {
	// создаем экземпляр WaitGroup
	wg := new(sync.WaitGroup)

	// увеличиваем счетчик WaitGroup на 1
	wg.Add(1)
	// запускаем горутину и передаем в нее ссылку на WaitGroup
	go worker1(wg)

	// ждем, когда обнулится счетчик WaitGroup
	wg.Wait()
}

func worker1(wg *sync.WaitGroup) {
	// 5 раз выводим сообщение
	for i := 0; i < 5; i++ {
		fmt.Println("Waiting...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Time has come!")
	// и уменьшаем счетчик WaitGroup
	wg.Done()
}
