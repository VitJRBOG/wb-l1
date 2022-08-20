package windows

import "fmt"

// Windows структура, олицетворяющая компьютер на Windows
type Windows struct{}

// InsertIntoUSBPort метод, реализующий подключение через USB
func (w Windows) InsertIntoUSBPort() {
	fmt.Println("USB connector is plugged into Windows machine.")
}
