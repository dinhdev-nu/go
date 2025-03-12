package main

import (
	"encoding/json"
	"fmt"
)

var jsonString = `{"name": "gopher", "age": 10}`

func main() {
	var result map[string]interface{}

	json.Unmarshal([]byte(jsonString), &result)

	convertJson, _:= json.Marshal(result)
	newJson := string(convertJson)
	
	fmt.Println(newJson)

}

// import (
// 	"fmt"
// 	"reflect"
// )

// func printType(i interface{}) {
// 	fmt.Println("Kiểu dữ liệu:", reflect.TypeOf(i))
// }

// func main() {
// 	printType(42)         // Output: int
// 	printType("Hello")    // Output: string
// 	printType([]int{1, 2}) // Output: []int
// }

// type Animal interface {
// 	Speak() string
// }

// type Dog struct {
// 	name string
// }

// type Cat struct {
// 	name string
// }

// func (d Dog) Speak() string {
// 	return "Gau Gau"
// }

// func (c Cat) Speak() string {
// 	return "Meo Meo"
// }

// func AnimalSpeak(a Animal) {
// 	fmt.Println(a.Speak())
// }

// func main() {
// 	cat := Cat{"Tom"}
// 	dog := Dog{"Jerry"}

// 	AnimalSpeak(cat)
// 	AnimalSpeak(dog)
// }

// interface{} là kiểu dữ liệu rỗng, nó có thể chứa giá trị của bất kỳ kiểu dữ liệu nào.
// interface{} trong go là 1 tập các method có thể implement bởi bất kỳ kiểu dữ liệu nào.

// not inhertiance(kế thừa) in go