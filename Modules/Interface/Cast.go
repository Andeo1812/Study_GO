package Interface

import "fmt"

func BuyTypes(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Nalic")
	case *Card:
		plasticCard, ok := p.(*Card)

		if !ok {
			fmt.Println("error conversion to type Card")
		}

		fmt.Println("Insert card,", plasticCard.Holder)
	default:
		fmt.Println("Something new!")
	}

	Buy(p)
}

// Также есть конструкция func ByuTypes(in interface{}), в которой приходится проверять тип самостоятельно в runtime
func ByuSomeInterface(in interface{}) {
	var p Payer
	var ok bool

	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T is not pay tool\n\n", in)

		return
	}

	BuyTypes(p)
}

func Cast() {
	fmt.Println("Cast")

	myWall := &Wallet{Cash: 10}
	BuyTypes(myWall)

	myCard := &Card{Balance: 0, Holder: "Oleg"}
	BuyTypes(myCard)

	ByuSomeInterface([]int{1, 2, 3})
}
