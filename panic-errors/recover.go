package panicerrors

import (
	"errors"
	"fmt"
)

func ThrowPanic(exrp int) (err error) {

	defer func() {

		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("a panic occurred but it is ok")
		} else if p != nil {
			panic("an unexpected error occurred and we do not want to recover")
		}
	}()

	switch exrp {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		panic(errCatastrophicReader) //it can be handle
	case 4:
		panic("unexpected error")
	case 5:
		err := (errors.New("ERROR"))
		return err
	default:
		fmt.Println("default")
	}

	return err
}

var errCatastrophicReader = errors.New("something catastrophic occurred in the reader")
