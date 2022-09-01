package BasicFunctional

import "fmt"

func Control() {
	fmt.Println("Control")

	//  Простое условие

	boolValue := true

	if boolValue {
		fmt.Println("boolValue -", boolValue)
	}

	mapValue := map[string]string{"name": "Oleg"}

	//  Условие с блоком инициализации
	if keyValue, keyExist := mapValue["name"]; keyExist {
		fmt.Println("name -", keyValue)
	}

	//  Получаем только признак существования ключа
	if _, keyExist := mapValue["name"]; keyExist {
		fmt.Println("key 'name' - exist")
	}

	//  Множественные if else
	cond := 1
	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2 {
		fmt.Println("cond is 2")
	}

	//  switch по 1 переменной
	strValue := "name"
	switch strValue {
	case "name":
		fmt.Println("name")
		fallthrough
	case "test", "lastName":
		fmt.Println("test", "lastName")
	default:
		fmt.Println("default")
	}

	// switch  как замена многим if else
	var value1, value2 = 2, 2
	switch {
	case value1 > 1 || value2 < 11:
		fmt.Println("First case")
	case value2 > 100:
		fmt.Println("Second case")
	}

	//  Выход из цикла, находясь в switch
Loop:
	for key, val := range mapValue {
		fmt.Println("switch in loop", key, val)
		switch {
		case key == "lastName":
			break
			fmt.Println("dont pront this")
		case key == "firstName" && val == "Vasily":
			fmt.Println("switch - break loop here")
			break Loop
		}
	} //  Конец for
}
