package main

import "fmt"

type Cat struct {
}

func (*Cat) Meow() {
	fmt.Println("Meow meow")
}

type Speaker interface {
	Speak()
}

type Animal struct {
	*Cat
}

func (adapter *Animal) Speak() {
	adapter.Meow()
}

func NewAdapter(cat *Cat) Speaker {
	return &Animal{cat}
}

func main() {
	animal := NewAdapter(&Cat{})
	animal.Speak()
}
