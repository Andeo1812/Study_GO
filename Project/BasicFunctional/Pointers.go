package BasicFunctional

import "fmt"

func Pointers() {
	fmt.Println("Pointers")

	a := 2
	b := &a
	*b = 3  //  a = 3
	c := &a //  новый указатель на переменную a

	//  Получение указателя на переменную типа int
	//  Инициализированной значение по умолчанию
	d := new(int)
	*d = 12
	*c = *d //  c = 12 -> a = 12
	*d = 13 //  c и a не изменились

	c = d   //  теперь c указывает туда же, куда и d
	*c = 14 //  c = 14 -> d = 14, a = 12
}
