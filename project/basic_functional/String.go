package basic_functional

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	fmt.Println("String")

	//  Пустая строка по-умолчанию
	var str string

	//  Со спец символами
	var hello string = "Hello\n\t"

	//  Без спец символов, \t тоже печатаются
	var world string = `World\n\t
	ipasfipafs
	pioajsfpjafs
	poajsfojasf
	japosfj
	`

	fmt.Println("str", str)
	fmt.Println("hello", hello)
	fmt.Println("world", world)

	var helloWorld = "Hello, world!"
	symbol := "象征"

	fmt.Println("helloWorld", helloWorld)
	fmt.Println("symbol", symbol)

	//  Одинарные кавычки для байт (uint8)
	var rawBinary byte = '\x27'

	//  rune (uint32) для UTF-8 символов
	var someChinese rune = '征'

	fmt.Println(rawBinary, someChinese)

	helloWorld = "Hello World 🙂"
	//  Конкатезация строк
	andGoodMorning := helloWorld + " and good morning!"

	fmt.Println(helloWorld, andGoodMorning)

	//  Строки неизменяемы
	//  cannot assign to helloWorld[0]
	//  helloWorld[0] = 72

	//  Получение длинны строки
	byteLen := len(helloWorld)
	symbols := utf8.RuneCountInString(helloWorld)

	fmt.Println(byteLen, symbols)

	//  Получение подстроки, в байтах, не символах
	hello = helloWorld[:12]
	H := helloWorld[0]

	fmt.Println(hello, H)

	//  Конвертация в слайс байт и обратно
	byteString := []byte(helloWorld)
	helloWorld = string(byteString)

	fmt.Println(byteString, helloWorld)
}
