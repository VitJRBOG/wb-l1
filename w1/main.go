package main

import "fmt"

// Human базовая структура
type Human struct {
	Name   string
	IsMale bool
	Age    int
	Height int
	Weight int
}

func (h *Human) Sit() {
	fmt.Printf("%s is sitting.\n", h.Name)
}

func (h *Human) Eat() {
	fmt.Printf("%s is eating.\n", h.Name)
}

func (h *Human) Stand() {
	fmt.Printf("%s is standing.\n", h.Name)
}

func (h *Human) Walk() {
	fmt.Printf("%s is walking.\n", h.Name)
}

func (h *Human) Lie() {
	fmt.Printf("%s is lying.\n", h.Name)
}

func (h *Human) Sleep() {
	fmt.Printf("%s is sleeping.\n", h.Name)
}

// Action структура, в которую встроена базовая структура Human
type Action struct {
	Human
}

func main() {
	a := Action{}
	a.Name = "Trish"
	a.IsMale = false
	a.Age = 25
	a.Height = 175
	a.Weight = 55

	a.Sit()
	a.Eat()
	a.Stand()
	a.Walk()
	a.Lie()
	a.Sleep()
}
