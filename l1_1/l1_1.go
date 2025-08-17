package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет\n", h.Name, h.Age)
}

type Action struct {
	Human
	Skill string
}

func (a Action) DoSomething() {
	fmt.Printf("%s умеет %s\n", a.Name, a.Skill)
}

func main() {
	a := Action{
		Human: Human{
			Name: "Виктория",
			Age:  27,
		},
		Skill: "программировать",
	}

	a.SayHello()
	a.DoSomething()
}
