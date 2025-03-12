package main

import "fmt"


type Node struct {
	data int
}
type Quantity struct {
	data int
}

type Book struct {
	title    string
	author   string
	price    float32
	quantity *Quantity
}

func main() {
	// Test()
	// Init()	
	// Example
	book1 := &Book{title: "Go Programming", author: "John Doe", price: 25.50, quantity: NewQuantity(10)}

	fmt.Printf("Book Name %s, Author %s, Price %.2f, Quantity %d\n", book1.title, book1.author, book1.price, book1.quantity.data)

	book1.changePrice(30.50)

	fmt.Printf("Book Name %s, Author %s, Price %.2f, Quantity %d\n", book1.title, book1.author, book1.price, book1.quantity.data)

	book2 := book1

	fmt.Printf("Book Name %s, Author %s, Price %.2f, Quantity %d\n", book2.title, book2.author, book2.price, book2.quantity.data)

}

func NewPointer() {
	pnt:= func () *int {
		p:= 10
		return &p
	}()

	fmt.Println(pnt) // "0xc0000b6010"
	fmt.Println(*pnt) // "10"
}

func Test() {
	var i *int; *i = 10
	fmt.Println(i) // "0xc0000b6010"
	fmt.Println(*i) // "10"


	func (i * int)  {
		*i = 20
		
	}(i)
	fmt.Println(*i) // "20"

	var j int = 20 

	func (j *int)  {
		*i = 30
	}(&j)
	fmt.Println(j) // "30"
}

func Init(){
	// x := 10
	// Change(&x)
	// fmt.Println(x) // "20"
	// p := &x // p, of type *int, points to x

	// fmt.Println(*p) // "10"
	// fmt.Println(p) // "0xc0000b6010"

	// *p = 20
	// fmt.Println(x) // "20"
	// fmt.Println(*p) // "20"
	// fmt.Println(p) // "0xc0000b6010"
	a:= &Node{data: 10} // == var a *Node; a = &Node{data: 10}
	fmt.Println(&a) // "0xc0000b6010"
	fmt.Println(a)
	ChangeData(a)
	fmt.Println(a)
	b:= &Node{data: 20}
	fmt.Println(&b)
	fmt.Println(a)
	fmt.Println(b)
}

func (b *Book) changePrice(price float32) {
	b.price = price
}
// contructer
func NewQuantity(data int) *Quantity {
	return &Quantity{data: data}
}
func ChangeData(a *Node) {
	a.data = 20
}

func Change(a *int) {
	*a = 20
}