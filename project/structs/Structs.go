package structs

import "fmt"

type Person struct {
	id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Person
}

func Structs() {
	fmt.Println("structs")

	//  Полное объявление
	var acc Account = Account{
		Id:   1,
		Name: "Oleg",
		Person: Person{
			id:      2,
			Name:    "Vasa",
			Address: "Kirgistan",
		},
	}

	fmt.Printf("%#v\n", acc)

	//  Короткое объявление. В прямом порядке без пропусков
	acc.Person = Person{2, "Roma", "Moscow"}

	fmt.Printf("%#v\n", acc)

	fmt.Println(acc.Name)
}
