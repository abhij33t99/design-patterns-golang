package behavioral

type Department interface {
	execute(*Patient)
	setNext(Department)
}
