package behavioral

import "fmt"

type Observer interface {
	update(string)
	getId() string
}

type Customer struct {
	id string
}

func (c *Customer) getId() string {
	return c.id
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s with item %s\n", c.id, itemName)
}
