package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

func main() {

	// Variable()
	// Loop()
	// Condition()
	ExeciseArray()

}

// Function
func Void () {}
func Add (a int, b int) int { return a+b }
func Change (a *int) { *a = 10 }
func MultipleReturn (a int, b int) (int, int) { return a+b, a-b }


func Condition () {
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error", err)
	}

	// if else
	if a > 10 {
		fmt.Println("Greater than 10")
	} else if a < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Equal to 10")
	}

	// switch case
	switch a {
		case 11:
			fmt.Println("Greater than 10")
		case 9:
			fmt.Println("Less than 10")
		default:
			fmt.Println("Equal to 10")
	}

}

func Loop () {
	arr := [3]int{1,2,3}
	dict := map[string]int{"a":1, "b":2};

	// for loop
	for i:= 0; i<3; i++ {
		fmt.Println("Index", i+1,arr[i])
	}

	// for loop with range
	for key, val := range dict {
		fmt.Printf("Key %s, Value %d\n", key, val)
	}

	// while loop
	i := 0
	for i < 3 {
		fmt.Println("Index", i+1,arr[i])
		i++
	}

}

func Variable () {
	// int 
	var a int = 10
	a1:= 10

	// float 
	var b float64 = 10.10
	// b1:= 10.10

	// string 
	var c string = "Hello"
	// c1:= "Hello"
	// c2 := fmt.Sprintf("Hello %s", "World")

	// boolean 
	var d bool = true
	// d1:= true

	// array 
	var arr = [3]int{1,2,3}

	// dictionary
	var dict = map[string]int{"a":1, "b":2};

	// fmt 
	fmt.Print("hlo") // print without new line
	fmt.Printf("Number %d, Float %f, %.2f, String %s, Boolean %t", a, b, b, c, d) // print with format
	fmt.Println(a, a1)

	fmt.Println(arr)
	fmt.Println(dict)

}

func HandleString() {

	// strings package to handle string

	var a string = "Hello"
	b := strings.ToTitle(a)
	fmt.Println(b) // "HELLO"
}
func HandleArray() {

	// slices package to handle array
	var a = []int{1,2,3,4,5}

	var newCopy = make([]int, len(a))
	copy(newCopy, a)
	fmt.Println(newCopy) // "[1 2 3 4 5]" 


	fmt.Println(cap(a)) // "5"
	c := a[1:3]
	fmt.Println(c) // "[2 3]"
	fmt.Printf("%d", cap(c)) // "4"

	// vd 
	c[0] = 10
	fmt.Println(a) // "[1 10 3 4 5]"

	// make 
	d := make([]int, 5, 10) // 5 is length, 10 is capacity
	fmt.Println(d) // "[0 0 0 0 0]"
}

func ExeciseArray() {
	a := []int{}
	for i:=0; i<10; i++ {
		a = append(a, i)
	}
	fmt.Println(a) // "[0 1 2 3 4 5 6 7 8 9]"

		index:=5;
		a = append(a[:index], a[index+1:]...)
		fmt.Println("Remove::: ", a) // "[0 1 2 3 4 6 7 8 9]"
	 
		max := slices.Max(a)
		fmt.Println("Max::: ", max) // "9"

		// reverse
		for i, j := 0, len(a)-1; i<j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		fmt.Println("Reverse::: ", a) // "[9 8 7 6 4 3 2 1 0]"
		// slices.Reverse(a)
		// fmt.Println(a) // "[9 8 7 6 4 3 2 1 0]"

		TestArray(a)

		fmt.Println(a) // "[1000 8 7 6 4 3 2 1 0]"


}

// slice and map are reference type so it will change the original value
func TestArray(a []int) {
	a[0] = 1000
}

func HandleMap() {
	// maps package to handle map
	var a = map[string]int{"a":1, "b":2}
	b:= maps.Keys(a)
	fmt.Println(b) // "[a b]"
	a["c"] = 3
	fmt.Println(a) // "map[a:1 b:2 c:3]"

	delete(a, "c")

	var d = make(map[string]int)
	d["a"] = 1

	// thao tac voi map
	//check exists
	if val, ok := a["a"]; ok {
		delete(a, "a")	
		fmt.Println(val)
	} else {
		fmt.Println("Not exists")
	}

	newMap := a
	newMap["a"] = 4
	fmt.Println(a) // "map[a:4 b:2]"


	copyMap := maps.Clone(a)
	fmt.Println(copyMap) // "map[a:1 b:2]"
	copyMap["a"] = 4
	fmt.Println(a) // "map[a:1 b:2]"
	fmt.Println(copyMap) // "map[a:4 b:2]"

	mapSlice := map[string]map[string][]int{
		"a": {"a": {1,2,3}},
		"b": {"b": {4,5,6}},
	}
	fmt.Println(mapSlice) // "map[a:[1 2 3] b:[4 5 6]]"

}

func execiseMap() {
	mapac := make([]interface{}, 0)
	mapac = append(mapac, map[string]int{"a":1})
	mapac = append(mapac, map[string]int{"b":2})
	fmt.Println(mapac) // "[map[a:1] map[b:2]]"

	
}

// Viết một hàm kiểm tra Map có lồng Map không
func CheckMap(mapSlice map[string]interface{}) bool {
	for _, val := range mapSlice {
		if _, ok := val.(map[string]interface{}); ok {
			return true
		}
	}
	return false
}

func Handle() {
	
	// number
	var a = "10"
	b, _ := strconv.ParseInt(a, 10, 64)
	fmt.Println(b) // "10"
	
}