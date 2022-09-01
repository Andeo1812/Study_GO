package Interface

import "fmt"

type Card struct {
	Balance int
	Holder  string
	CVV     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("not such money")
	}

	c.Balance -= amount

	return nil
}

func Many() {
	fmt.Println("Many")

	myWall := &Wallet{Cash: 10}
	Buy(myWall)

	myPhone := &Phone{Money: 10}
	Buy(myPhone)

	myCard := &Card{Balance: 10, Holder: "Oleg"}
	Buy(myCard)
}
