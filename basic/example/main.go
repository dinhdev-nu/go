package main

import (
	"fmt"
	"reflect"
)


type Car struct {
	price int
}

func main() {
	// a := 10
	// b := 20
	// Swap(&a, &b)
	// fmt.Printf("Valu a: %d, Value b: %d\n", a, b)

	// bai1()

	// i:= 10

	// fmt.Print(reflect.TypeOf(i))
	x := 3.14
	v := reflect.ValueOf(x) // Lấy giá trị của x
	fmt.Println("Giá trị:", v)
	fmt.Println("Giá trị kiểu float:", v.Float()) // Lấy giá trị float

	// Type Assertion
	// use .Int() , .Uint() , .Float() , .Bool() , .Complex(), ... to get value of reflect.Value


	v1:= reflect.ValueOf(&x).Elem()
	v1.SetFloat(3.1415)
	fmt.Println("Giá trị mới của x:", x)
}

func bai1() {
	car1:= new(Car)
	car1.price = 1000
	fmt.Println(car1)
	car1.updateCart(2000)
	fmt.Println(car1)
	changePrice(car1, 3000)
	fmt.Println(car1)
}

func (c *Car) updateCart (price int) {
	c.price = price
}

func changePrice(c *Car, p int){
	c.price = p
}

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}