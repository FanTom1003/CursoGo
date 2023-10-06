package src

import "fmt"

func ciclos() {
	//ciclo for
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	value := 0
	//for while
	for value < 10 {
		fmt.Println(value)

		value++
	}
	value2 := 0
	//for forever
	for {
		fmt.Println(value2)
		value2++
	}
}
