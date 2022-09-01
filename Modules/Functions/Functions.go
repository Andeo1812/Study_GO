package Functions

import "fmt"

// Обычное определение
func singleIn(in int) int {
	return in
}

// Много парметров
func multiIn(a, b int, c int) int {
	return a * b * c
}

//  Именованный результат
func namedReturn() (out int) {
	//  out = 0
	out = 2

	return
}

//  Несколько результатов
func mulriReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}

	return in, nil
}

//  Несколько именованных резульатов
func multiNamedReturn(ok bool) (res int, err error) {
	res = 1
	if ok {
		err = fmt.Errorf("some error happend 2")
		//  Аналогично return

		return 3, fmt.Errorf("some error happend 2")
	}

	res = 2

	return
}

func sum(in ...int) (result int) {
	fmt.Printf("in := %#v\n", in)
	for _, val := range in {
		result += val
	}

	return
}

func Functions() {
	fmt.Println("Functions")

	sum(1, 23, 44, 55, 101)

	nums := []int{1, 2, 3, 4}
	fmt.Println(nums, sum(nums...))
}
