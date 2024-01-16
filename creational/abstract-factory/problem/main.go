package main

import "fmt"

type Drink interface {
	Drink()
}
type Food interface {
	Eat()
}

type Voucher struct {
	Drink
	Food
}

type Coffee struct{}

func (Coffee) Drink() {
	fmt.Printf("It's coffee, drinkable")
}

type Beer struct{}

func (Beer) Drink() {
	fmt.Printf("It's beer, drinkable")
}

type Cake struct{}

func (Cake) Eat() {
	fmt.Println("It's cake, eatable")
}

type GrilledOctoupus struct{}

func (GrilledOctoupus) Eat() {
	fmt.Println("It's Grilled Octopus, eatable")
}

func main() {
	fmt.Println([]Voucher{
		// Voucher with Coffee & Cake, it's good
		Voucher{
			Drink: Coffee{},
			Food:  Cake{},
		},
		// This voucher is great!!
		Voucher{
			Drink: Beer{},
			Food:  GrilledOctoupus{},
		},
		// Oh, this voucher quite weird, I'm not sure if I can use it
		Voucher{
			Drink: Coffee{},
			Food:  GrilledOctoupus{},
		},
	})
}
