/*
 *   Copyright (c) 2021 Anton Brekhov
 *   All rights reserved.
 */
package composer

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Composer struct {
	Name     string
	InitChan int
}

type ChildOfComposer struct {
	Composer
}

type Composable interface {
	Compose()
}

// Checks!

func CheckRun() {
	initChan := make(chan int)

	c := New("Mather")
	outChanValue := Compose(c, reflect.ValueOf(initChan), "M10", "M2", "D5")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		integer, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		initChan <- integer
		output, ok := outChanValue.Recv()
		fmt.Printf("ok: %v\n", ok)
		fmt.Printf("output: %v\n", output)
		fmt.Printf("len(initChan): %v\n", len(initChan))
		fmt.Printf("outChanValue.Len(): %v\n", outChanValue.Len())
	}

}

func CheckOneNum(x int) (y int) {
	initChan := make(chan int)

	c := New("Mather")
	outChanValue := Compose(c, reflect.ValueOf(initChan), "M10", "M2", "D5")
	initChan <- x
	output, _ := outChanValue.Recv()
	y = output.Interface().(int)
	return y
}

func CheckChildOneNum(x int) (y int) {
	initChan := make(chan int, 100)

	child := &ChildOfComposer{}
	outChanValue := Compose(child, reflect.ValueOf(initChan), "M10", "M2", "D5")
	initChan <- x
	output, _ := outChanValue.Recv()
	y = output.Interface().(int)
	return y
}

// COMPOSER
func New(name string) *Composer {
	return &Composer{Name: name}
}

func Inputer(intergers ...int) chan int {
	initChan := make(chan int)
	for _, i := range intergers {
		initChan <- i
	}
	return initChan
}

func (*Composer) M10(in chan int, out chan int) {
	for i := range in {
		out <- i * 10
	}
}

func (*Composer) M2(in chan int, out chan int) {
	for i := range in {
		out <- i * 2
	}
}

func (*Composer) D5(in chan int, out chan int) {
	for i := range in {
		out <- i / 5
	}
}

// CHILDOFCOMPOSER

func (*ChildOfComposer) M10(in chan int, out chan int) {
	for i := range in {
		out <- i * 10
	}
}

func (*ChildOfComposer) M2(in chan int, out chan int) {
	for i := range in {
		out <- i * 2
	}
}

func (*ChildOfComposer) D5(in chan int, out chan int) {
	for i := range in {
		out <- i / 5
	}
}

// https://play.golang.org/p/WJR6tRIpPZ
// https://golang.hotexamples.com/examples/reflect/-/MakeChan/golang-makechan-function-examples.html
// https://imatmati.github.io/posts/golang-reflection

func Compose(i interface{}, initChan reflect.Value, fns ...string) reflect.Value {
	var outChanValue reflect.Value
	var prevChanValue, nextChanValue reflect.Value
	for fninx, fn := range fns {
		fnx := reflect.ValueOf(i).MethodByName(fn)
		if !fnx.IsValid() {
			fmt.Println("Method not valid")
		}
		if fnx.IsNil() {
			fmt.Println("Method is nil")
		}
		if fnx.IsZero() {
			fmt.Println("Method is zero")
		}
		t := fnx.Type()
		var argList []reflect.Value
		for j := 0; j < t.NumIn(); j++ {
			if fninx == 0 && j == 0 {
				// v2 := reflect.ValueOf(initChan)
				v2 := initChan
				prevChanValue = v2
				argList = append(argList, v2)
				continue
			}

			if j == 0 { // input then prevChan
				argList = append(argList, prevChanValue)
			} else {
				t2 := t.In(j)
				v := reflect.MakeChan(t2, 1)

				nextChanValue = v
				argList = append(argList, nextChanValue)
				prevChanValue = nextChanValue
			}
			if fninx == len(fns)-1 && j == t.NumIn()-1 {
				outChanValue = nextChanValue
			}
		}
		go fnx.Call(argList)
	}
	return outChanValue
}
