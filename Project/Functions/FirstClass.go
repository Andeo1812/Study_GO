package Functions

import "fmt"

// Обычная функция
func something() {
	fmt.Println("regular func")
}

func FirstClass() {
	fmt.Println("FirstClass")

	//  Анонимная функция
	func(in string) {
		fmt.Println("anon func out:", in)
	}("nobody")

	//  Присваивание анонимной функции в переменную
	printer := func(in string) {
		fmt.Println("printer out:", in)
	}

	printer("as var")

	//  Определяем тип функции
	type strFuncType func(string)

	//  Функция принимает callback
	worker := func(callback strFuncType) {
		callback("as callback")
	}

	worker(printer)

	//   Функция возвращает замыкание
	prefixer := func(prefix string) func(string) {
		//  prefix
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}

	successLogger := prefixer("SUCCESS")
	successLogger("expected behavior")
}
