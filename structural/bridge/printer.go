package structural

import "fmt"

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by EPSON printer")
}

type Hp struct{}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by HP printer")
}
