package behavioral

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type PassengerTrain struct {
	mediator Mediator
}

func (t *PassengerTrain) arrive() {
	if !t.mediator.canArrive(t) {
		fmt.Println("Passenger Train can't arrive")
		return
	}
	fmt.Println("Passenger Train arriving")
}

func (t *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	t.mediator.notifyAboutDeparture()
}

func (t *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	t.arrive()
}

type FreightTrain struct {
	mediator Mediator
}

func (g *FreightTrain) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.arrive()
}
