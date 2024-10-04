package structural

import "fmt"

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct{
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}