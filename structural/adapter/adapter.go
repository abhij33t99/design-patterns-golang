package structural

import "fmt"

type WindowsAdapter struct {
	windows *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts lightning to usb to support on windows machine")
	w.windows.InsertIntoUSBPort()
}
