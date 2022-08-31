package basic_tools

import "fmt"

const pi float32 = 3.141

// Инициалзация в блоке
const (
	hello = "Hello"
	e     = 2.718
)

const (
	zero = iota //  0, iota - специальный итератор
	_           // placeholder - пропуск
	two
	three //  3
)

const (
	_         = iota             //  Пропускаем первое значение
	KB uint64 = 1 << (10 * iota) //  1 << (10 * 1) = 1024
	MB                           //  1 << (10 * 2) = 1048576
	GB
	TB
	PB
	EB
)

const (
	//  Нетипизированная константа
	year = 2022
	//  Типизированная константа
	yearTyped int = 2022
)

func Const() {
	fmt.Println("Const")

	var month int32 = 13

	//  тип year подстраивается в зависимости от места использования т.к. нетипизированный тип
	fmt.Println(month + year)

	//  month + yearTyped (mismatched types int32 and int)
	//  fmt.Println(month + yearTyped)
}
