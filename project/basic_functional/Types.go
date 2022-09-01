package basic_functional

import "fmt"

type UserID int

func Types() {
	fmt.Println("Types")

	idx := 1

	var uid UserID = 42

	//  Даже если базовый тип однознаковый, разные типы несовместимы
	//  cannot use uid (type UserID) as type int64 in assigment
	//  myID := inx

	myID := UserID(idx)
	int_ := int(myID)

	fmt.Println(uid, myID, int_)
}
