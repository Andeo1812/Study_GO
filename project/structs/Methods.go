package structs

import "fmt"

type Person_ struct {
	Id   int
	Name string
}

func SumAges(p1, p2 Person_) int {
	return p1.Id + p2.Id
}

// Не изменит оригинальной структуры, для которой вызван
func (p Person_) UpdateName(name string) {
	p.Name = name
}

// Изменяет оригинальную структуру
func (p *Person_) SetName(name string) {
	p.Name = name
}

type Account_ struct {
	Id   int
	Name string
	Person_
}

// Изменяет оригинальную структуру
func (p *Account_) SetName(name string) {
	p.Name = name
}

type MySplice []int

func (sl *MySplice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySplice) Count() int {
	return len(*sl)
}

func Methods() {
	fmt.Println("Methods")

	pers := Person_{1, "Vasa"}
	// pers := new(Person_)
	pers.SetName("Yasa")
	// (&pers).SetName("Yasa")
	//  fmt.Printf("updated person: %#v\n", pers)

	var acc Account_ = Account_{
		Id:   1,
		Name: "Rva",
		Person_: Person_{
			Id:   2,
			Name: "Yadas",
		},
	}

	acc.SetName("romaaa")
	acc.Person_.SetName("Tesatas")

	sl := MySplice([]int{1, 2})

	sl.Add(5)

	fmt.Println(sl.Count(), sl)

	fmt.Printf("%#v\n", acc)
}
