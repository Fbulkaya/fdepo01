package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}
type Book struct {
	desi int
}
type Electronic struct {
	desi int
}
type Flower struct {
	desi int
}
type Shoes struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}
func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*2
}
func (flower *Flower) ShippingCost() int {
	return 5 + flower.desi*2
}
func (shoes *Shoes) ShippingCost() int {
	return 5 + shoes.desi*2
}
func main() {
	var product IShippable

	product = &Book{desi: 10}
	fmt.Println(product.ShippingCost())
	product = &Electronic{desi: 20}
	fmt.Println(product.ShippingCost())
	product = &Flower{desi: 5}
	fmt.Println(product.ShippingCost())
	product = &Shoes{desi: 7}
	fmt.Println(product.ShippingCost())

}
