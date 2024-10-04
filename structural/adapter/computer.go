package structural

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Inserted into lightning port on mac machine")
}

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("Inserted into usb port on windows machine")
}
