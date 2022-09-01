package basic_functional

import "fmt"

func Map() {
	fmt.Println("Map")

	//  Инициализация при создании
	user := map[string]string{}

	user["test"] = "r"

	//  Инициализация с заданной емкостью
	profile := make(map[string]string, 10)

	//  Колличество элементов
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)

	//  Если ключа нет - вернет значение по умолчанию для типа
	mName := user["middleName"]

	fmt.Println("nName", mName)

	//  Проверка на существование ключа
	mName, mNameExist := user["middleName"]

	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	//  Пустая переменная - только проверяем что ключ есть
	_, mNameExist2 := user["middleName"]

	fmt.Println("mNameExist2:", mNameExist2)

	//  Удаление ключа
	delete(user, "lastName")

	fmt.Printf("%#v\n", user)
}
