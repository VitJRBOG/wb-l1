package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("I'm going sleep...")
	MySleepFunc(3 * time.Second)
	fmt.Println("I'm awake! Hello world!")
}

func MySleepFunc(waitFor time.Duration) {
	t := time.NewTimer(waitFor) // создаем новый таймер из пакета time
	<-t.C                       // ждем, когда канал в экземпляре таймера что-нибудь вернет через время waitFor
	t.Stop()                    // таймер больше не нужен, завершаем его
}
