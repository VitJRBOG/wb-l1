package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Number struct {
	N int
}

func main() {
	n := Number{} // объявляем структуру-счетчик

	// создаем экземпляр контекста с таймером, чтобы программа не работала вечно
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	// создаем экземпляр мьютекса, чтобы избежать гонки
	mutex := new(sync.Mutex)
	// создаем экземпляр WaitGroup, чтобы программа не завершилась раньше горутин
	wg := new(sync.WaitGroup)

	// запускаем горутины с инкрементами
	for i := 5; i >= 0; i-- {
		wg.Add(1)
		go count(ctx, mutex, wg, &n)
	}

	// ждем, пока горутины завершатся
	wg.Wait()
	fmt.Println(n.N)
}

func count(ctx context.Context, mutex *sync.Mutex, wg *sync.WaitGroup, n *Number) {
Loop: // описываем метку, чтобы можно было завершить цикл из блока select
	for {
		select {
		case <-ctx.Done(): // проверяем, не завершился ли контекст
			break Loop
		default:
			mutex.Lock() // блокируем мьютекс, чтобы не было гонки
			n.N++        // спокойно инкрементируем число
			fmt.Println(n.N)
			mutex.Unlock() // разблокируем мьютекс, чтобы другие горутины могли его заблокировать
		}
	}
	wg.Done() // обозначаем завершение горутины
}
