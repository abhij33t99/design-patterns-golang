package behavioral

import "fmt"

type Device interface {
	on()
	off()
}

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("TV is being turned on")
}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("TV is being turned off")
}