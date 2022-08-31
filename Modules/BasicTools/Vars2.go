package BasicTools

import "fmt"

func Vars2() {
	fmt.Println("Vars_2")
	//  int
	//  int - платформозависимый тип, 32/64
	var i int = 10

	//  Автоматический выбранный int
	var autoInt = -10

	//  int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// uint - платформозависимый тип, 32/64
	var unsignedInt uint = 100500

	//  uint8, uint16, uint32, uint64
	var unsignedBigInt uint64 = 1<<64 - 1

	//  Преобразование типов для группы 1 типа (родственными)
	p := 9
	o := int64(p)

	fmt.Println(o, i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	//  float
	//  float32, float64
	var pi float32 = 3.141

	var e = 2.718

	fmt.Printf("%T", e)

	goldenRatio := 1.618

	fmt.Println(pi, e, goldenRatio)

	//  bool
	//  false по-умолчанию
	var b bool
	var isOk bool = true
	var success = true
	cond := true

	fmt.Println(b, isOk, success, cond)

	//  complex
	//  complex64, complex128
	var c0 complex128 = -1.1 + 7.12i
	c1 := -1.1 + 7.12i

	fmt.Println(c0, c1)
}
