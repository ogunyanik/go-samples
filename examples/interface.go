package examples

import (
	"fmt"
)

func AppendFunc() {
	fmt.Println("a")
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	programmers = append(programmers, elliot)

	// numbers := []float64{1.1, 2.2, 3.3}
	// x, _ := Sum_Deneme(numbers)
	// fmt.Println(x)
	//If you want to append like this. You should implement all interface method to struct (look at 1. and 2. )
	// You can try delete the implementation "2", you will see complier scream at you
	//Only you get the compiler error use the append
}

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string { //  1
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int { // 2
	return 21
}
