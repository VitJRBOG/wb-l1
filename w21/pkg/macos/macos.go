package macos

import "fmt"

// MacOS структура, олицетворяющая компьютер на MacOS
type MacOS struct{}

// InsertIntoLightningPort метод, реализующий подключение через Lightning
func (m MacOS) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into MacOS machine.")
}
