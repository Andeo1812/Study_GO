package _interface

import "fmt"

type Phone struct {
	Money   int
	AppleId string
}

func (p *Phone) Pay(amount int) error {
	if p.Money < amount {
		return fmt.Errorf("not such money")
	}

	p.Money -= amount

	return nil
}

func (p *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("Entry phone")
	}

	return nil
}

type Ringer interface {
	Ring(string) error
}

type NFCPhone interface {
	Payer
	Ringer
}

func PayForMeWithPhone(phone NFCPhone) {
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("Error pay - %v\n\n", err)
		return
	}

	fmt.Printf("Success paying %T\n\n", phone)
}

func Embed() {
	fmt.Println("Embed")

	myPhone := &Phone{Money: 10}

	PayForMeWithPhone(myPhone)
}
