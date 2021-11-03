package main

import (
	"fmt"
	"time"

	// Channel "Users/ogun.yanik/Documents/GitHub/go-samples/GoRoutine-Channels"
	Channelw "Users/ogun.yanik/Documents/GitHub/go-samples/GoRoutine-Channels/SimpleChannel"

	// L "Users/ogun.yanik/Documents/GitHub/go-samples/concurrency"

	selectstatement "Users/ogun.yanik/Documents/GitHub/go-samples/GoRoutine-Channels/channel-controls/select-statement"
)

func main() {
	fmt.Println("a----------------" + time.Now().String())
	// mutex.Use()
	// forLoop.ChannelsControl()
	selectstatement.Test()
	// commaok.ChannelsControl()
	// channeltype.ChannelTypess()
	return
	// L.Handles()
	// Channel.Handles()
	Channelw.Calculate2()
	// This will throw an error
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)

	// var numbers = []float64{12.1, 22.3, 11.0}

	// abc, _ := Sum_Deneme(numbers)
	// fmt.Println(abc)

}

func callTheSum() {
	var numbers = []float64{12.1, 22.3, 11.0}

	_, _ = sum(numbers...)
}

func sum(values ...float64) (float64, error) { ///"values..." says that is array of something
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total, nil

}

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return 21
}
