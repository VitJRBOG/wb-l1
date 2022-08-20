package adapter

import (
	"fmt"
	"w21/pkg/windows"
)

// WindowsAdapter структура, олицетворяющая адаптер
type WindowsAdapter struct {
	WindowsPC windows.Windows
}

// InsertIntoLightningPort метод, реализующий подключение через Lightning
// у компьютера на Windows:
// внутри себя вызывает метод подключения через USB
func (w WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowsPC.InsertIntoUSBPort()
}
