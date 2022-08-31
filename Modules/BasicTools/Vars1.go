package BasicTools

import "fmt"

func Vars1() {
	fmt.Println("Vars_1")
	//  Значение по умолчанию. Для int = 0
	var num_0 int

	//  Значение при инициализации
	var num_1 int = 1

	//  Пропуск типа
	var num_2 = 20

	fmt.Printf("%d %d %d\n", num_0, num_1, num_2)

	//  Короткое объявление
	num__ := 30

	//  Перезапись переменной
	num__ = 50

	//  Инкрементация
	num__ += 1

	fmt.Println("+=", num__)

	// ++num не существует
	num__++

	fmt.Println("++", num__)

	//  Стиль camelCase - принят на GO. under_case - нет. IDE подсказывает убрать
	camelCase := 10
	under_case := 10

	fmt.Println(camelCase, under_case)

	// Объявление нескольких переменных
	var weight, height int = 10, 20

	//  Присвоение в существующие переменные
	weight, height = 11, 21
	weight, height = height, weight

	//  Короткое присваивание. Хотя бы 1 должна быть новой
	weight, age := 12, 22

	fmt.Println(weight, height, age)
}
