package cmd

import (
	"fmt"
	"w21/pkg/adapter"
	"w21/pkg/macos"
	"w21/pkg/windows"
)

// Computer - интерфейс, которому удовлетворяют все структуры,
// имеющие функцию подключения через Lightning
type Computer interface {
	InsertIntoLightningPort()
}

// User просто пользователь
type User struct{}

// InsertLightningConnectorIntoComputer принимает на вход интерфейс Computer,
// и вызывает у него функцию подключения через Lightning
func (u User) InsertLightningConnectorIntoComputer(c Computer) {
	fmt.Println("User inserts Lightning connector into computer.")
	c.InsertIntoLightningPort()
}

func Execute() {
	// создаем экземпляр пользователя
	u := User{}

	// создаем экземпляр компьютера на MacOS
	mac := macos.MacOS{}

	// подключить Lightning к компьютеру на MacOS пользователю не составляет труда,
	// так как есть совместимость
	u.InsertLightningConnectorIntoComputer(mac)

	// создаем экземпляр компьютера на Windows
	winPC := windows.Windows{}

	// так как у компьютера на Windows нет порта для подключения Lighning,
	// нужно создать экземпляр адаптера и передать в него экземпляр
	// компьютера на Windows
	a := adapter.WindowsAdapter{
		WindowsPC: winPC,
	}

	// а затем передать экземпляр адаптера на вход функции подключения через Lightning
	u.InsertLightningConnectorIntoComputer(a)
}
