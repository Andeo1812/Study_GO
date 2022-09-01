package basic_functional

import "fmt"

func Array() {
	fmt.Println("Array")

	//  Размер массива является частью его типа
	//  Инициализация значениями по-умолчанию
	var a1 [3]int //  [0, 0, 0]

	fmt.Println("array short", a1)
	fmt.Printf("array short %v\n", a1)
	fmt.Printf("array full  %#v\n", a1)

	//  Работает с константами
	const size = 2
	var a2 [2 * size]bool //  [false, false, false, false]

	fmt.Println("a2", a2)

	//  Определение размера при определении
	a3 := [...]int{1, 2, 3, 4, 5, 10}

	fmt.Println("a3", a3)

	//  Проверка при компиляции или при выполнение
	//  При статическом ind при выходе за длинну - ошибка компиляции
	//  При динамическом ind - invalid array index
	//  ind := 100500
	//  a3[ind] = 12
}
