package main

import (
	"fmt"

	"project3.com/test/child"
)

type Valueer interface {
	printf()
}

type Value1 struct {
	num int
}

type Value2 struct {
	num int
}

func (self *Value1) printf() {
	fmt.Println(self.num)
}

func (self *Value2) printf() {
	fmt.Println(self.num)
}
func main() {

	// var te1 *Valueer
	var pxee child.Value3
	pxee.Printf()

	p1 := new(Value1)
	p1.printf()

	var te1 Valueer
	te1 = p1
	// te1 = *p1
	te1.printf()

	var te2 Valueer
	te2 = &Value2{1}
	te2.printf()
}
