package Interface

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("no money")
	}

	w.Cash -= amount

	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Error paying %T: %v\n\n", p, err)
	}
	fmt.Printf("thancks\n")
}

func Basic() {
	fmt.Println("Basic")

	myWal := &Wallet{Cash: 100}

	Buy(myWal)
}
