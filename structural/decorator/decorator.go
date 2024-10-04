package structural

import "fmt"

type CheeseTopping struct {
	Pizza IPizza
}

func (c *CheeseTopping) GetPrice() int {
	return c.Pizza.GetPrice() + 5
}

type MozarelloTopping struct {
	Pizza IPizza
}

func (m *MozarelloTopping) GetPrice() int {
	return m.Pizza.GetPrice() + 10
}

func main() {
	pizza := NonVegPizza{}

	pizzaWithCheese := CheeseTopping{
		Pizza: &pizza,
	}

	pizzaWithMozarello := MozarelloTopping{
		Pizza: &pizzaWithCheese,
	}

	fmt.Println("Total price: ", pizzaWithMozarello.GetPrice())
}
