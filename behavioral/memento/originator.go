package behavioral

type Originator struct {
	state string
}

func (o *Originator) createMemento() *Memento {
	return &Memento{
		state: o.state,
	}
}

func (o *Originator) restoreMemento(m *Memento) {
	o.state = m.getSavedState()
}

func (o *Originator) getState() string{
	return o.state
}
