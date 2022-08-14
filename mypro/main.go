package main

import (
	"fmt"
	"mypro/user"
)

func main() {
	/* 	var name = "int" // c初始化變量
	   	var age = 132    // l類型推斷
	   	fmt.Printf("age: %v\n", age)
	   	s := user.Hello()
	   	fmt.Printf("s: %v\n", s)
	   	fmt.Printf("name: %v\n", name)
	   	num := 12369
	   	fmt.Printf("num: %v\n", num)
	   	const ch = 6
	   	fmt.Printf("ch: %v\n", ch)
	   	const (
	   		start = iota
	   		a     = iota
	   		b     = iota
	   		c     = iota
	   	)
	   	fmt.Printf("a: %v\n", a)
	   	fmt.Printf("b: %v\n", b)
	   	fmt.Printf("c: %v\n", c) */
	/* s := user.Hello()
	fmt.Printf("s: %v\n", s) */
	/* 	if true {
		return true
	} */
	s := user.Hello()
	fmt.Printf("s: %v\n", s)
	var name string = "json"
	age := 99
	b := true
	fmt.Printf("age: %v\n", age)
	fmt.Printf("%T\n", name)
	fmt.Printf("b: %v\n", b)
	p := &age
	fmt.Printf("p: %v\n", p)
	fmt.Printf("%T\n", p)
	a1 := [3]int{1, 22, 333}
	fmt.Printf("%T\n", a1)
	fmt.Printf("a1: %v\n", a1)
}
