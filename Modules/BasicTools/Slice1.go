package BasicTools

import "fmt"

func Slice1() {
	fmt.Println("Slice1")

	//  Инициализация
	var buf0 []int             //  length-0		capacity-0
	buf1 := []int{}            //  length-0		capacity-0
	buf2 := []int{32}          //  length-1		capacity-1
	buf3 := make([]int, 0)     //  length-0		capacity-0
	buf4 := make([]int, 5)     //  length-5		capacity-5
	buf5 := make([]int, 5, 10) //  length-5		capacity-10

	fmt.Println(buf0, buf1, buf2, buf3, buf4, buf5)

	//  Обращение к элементам
	someInt := buf2[0]

	//  Ошибка при выполнение
	//  panic: runtime error: index out of range
	//  someWrongInt := buf2[1]

	fmt.Println(someInt)

	//  Добавление элементов
	var buf []int            //  length-0		capacity-0
	buf = append(buf, 9, 10) //  length-2		capacity-2
	buf = append(buf, 12)    //  length-3		capacity-4

	//  Добавление другого слайса
	otherBuf := []int{0, 0, 0}     // [0, 0, 0]
	buf = append(buf, otherBuf...) //  length-6		capacity-8
	//  buf = append(buf, otherBuf[0], otherBuf[1], otherBuf[2])

	fmt.Println(buf, otherBuf)

	//  Информация о сплайсу
	var bufLen, bufCap int = len(buf), cap(buf)

	fmt.Println(bufLen, bufCap)

}
