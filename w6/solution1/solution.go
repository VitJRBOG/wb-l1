package solution1

import (
	"context"
	"fmt"
	"time"
)

func Execute() {
	// создаем экземпляр контекста с таймаутом в 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// запускаем горутину и передаем ей контекст
	go worker(ctx)

	// чтобы главная горутина не завершилась раньше дочерней, ждем завершения контекста и тут тоже
	<-ctx.Done()
}

func worker(ctx context.Context) {
	fmt.Println("I was born!")
	for {
		select {
		// проверяем, если контекст завершился, то вызываем return и завершаем горутину
		case <-ctx.Done():
			fmt.Println("I'm going to die...")
			return
		default:
			fmt.Println("I'm still alive!")
			time.Sleep(1 * time.Second)
		}
	}
}
