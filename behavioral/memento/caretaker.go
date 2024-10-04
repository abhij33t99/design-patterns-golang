package behavioral

type Caretaker struct {
	mementoList []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementoList[index]
}
