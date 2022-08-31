package BasicTools

import "fmt"

func Slice2() {
	fmt.Println("Slice2")

	buf := []int{1, 2, 3, 4, 5}

	//  Получение среза, указывающего на ту же память
	sl1 := buf[1:4] //  [2, 3, 4]
	sl2 := buf[2:]  //  [1, 2]
	sl3 := buf[:2]  //  [3, 4, 6]

	fmt.Println(sl1, sl2, sl3)

	//  Копирование слайса, получение области
	newBuf := buf[:] //  [1, 2, 3, 4, 5]
	newBuf[0] = 9    //  [9, 2, 3, 4, 5], та же память

	//  len int
	//  cap int
	//  data uintptr

	fmt.Println("newBuf cap", cap(newBuf))

	//  newBuf теперь указывает на другие данные
	newBuf = append(newBuf, 6)

	fmt.Println("new newBuf cap", cap(newBuf))

	newBuf[0] = 1
	//  buf 		= [9, 2, 3, 4, 5], не изменился
	//  newBuf 		= [1, 2, 3, 4, 5, 6], изменился

	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	//  Копирование одного слайса в другой
	var emptyBuf []int //  length-0		capacity-0
	//  Неправильно - скопирует меньшее (по length) из 2-х сплайсов
	copied := copy(emptyBuf, buf) //  copied = 0

	fmt.Println(copied, emptyBuf)

	//  Правильно
	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)

	fmt.Println(newBuf)

	// Можно копировать в часть существеющего слайса
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6}) //  ints = [1, 5, 6, 4]

	fmt.Println(ints)
}
