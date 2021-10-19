package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
	Play()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) Eat() {
	fmt.Printf("%s 吃饭", d.name)
}
func (d Dog) Sleep() {
	fmt.Printf("%s 睡觉", d.name)
}
func (d Dog) Play() {
	fmt.Printf("%s 玩", d.name)
}

func (c Cat) Eat() {
	fmt.Printf("%s 吃饭", c.name)
}
func (c Cat) Sleep() {
	fmt.Printf("%s 睡觉", c.name)
}
func (c Cat) Play() {
	fmt.Printf("%s 玩", c.name)
}

func main() {
	var a Animal
	var dog = Dog{name: "xiaogou"}
	var cat = Cat{name: "xiaomao"}
	a = dog
	a.Eat()
	a = cat
	a.Sleep()
}
