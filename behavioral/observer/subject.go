package behavioral

import "fmt"

type Subject interface {
	registerObserver(observer Observer)
	removeObserver(observer Observer)
	notifyObservers()
}

//concrete subject

type Item struct {
	observers []Observer
	name      string
	isStock   bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) registerObserver(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) removeObserver(observer Observer) {
	length := len(i.observers)
	for ind, o := range i.observers {
		if o.getId() == observer.getId() {
			i.observers[ind] = i.observers[length-1]
			i.observers = i.observers[:length-1]
			break
		}
	}
}

func (i *Item) notifyObservers() {
	for _, o := range i.observers {
		o.update(i.name)
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is in stock\n", i.name)
	i.isStock = true
	i.notifyObservers()
}
