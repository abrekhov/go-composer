/*
 *   Copyright (c) 2021 Anton Brekhov
 *   All rights reserved.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// InputData int
// MultiplierByTwo int int
// MultiplierByTen int int
// DeviderByFive int int
// CheckerOnEven int bool

type Composer struct {
	Name     string
	InitChan int
}

func main() {
	initChan := make(chan int)

	c := New("Mather")
	outChanValue := c.Compose(initChan, "M10", "M2", "D5")
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

// https://play.golang.org/p/WJR6tRIpPZ
// https://golang.hotexamples.com/examples/reflect/-/MakeChan/golang-makechan-function-examples.html
// https://imatmati.github.io/posts/golang-reflection

func (c *Composer) Compose(initChan chan int, fns ...string) reflect.Value {
	var outChanValue reflect.Value
	var prevChanValue, nextChanValue reflect.Value
	for fninx, fn := range fns {
		fnx := reflect.ValueOf(c).MethodByName(fn)
		t := fnx.Type()
		var argList []reflect.Value
		for j := 0; j < t.NumIn(); j++ {
			if fninx == 0 && j == 0 {
				v2 := reflect.ValueOf(initChan)
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
