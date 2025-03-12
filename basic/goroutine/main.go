// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type Order struct {
// 	message string
// }

// func main() {
// 	defer fmt.Println("Done")

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	messChanel := make(chan Order)

// 	orders := []Order{
// 		{message: "Hello"},
// 		{message: "World"},
// 		{message: "From"},
// 		{message: "Go"},
// 	}

// 	go Pub(messChanel, orders, &wg)
// 	go Sub(messChanel, &wg)

// 	wg.Wait()
// }

// func Pub (mess chan<- Order, orders []Order, sg *sync.WaitGroup){
// 	defer sg.Done()
// 	defer close(mess)
// 	for _, order := range orders {
// 		fmt.Println("Sending::: ", order.message)
// 		mess <- order
// 		time.Sleep(time.Second * 1)
// 	}

// }

// func Sub (mess <-chan Order, sg *sync.WaitGroup){
// 	defer sg.Done()
// 	for	order := range mess {
// 		fmt.Println("Message::: ", order.message)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func SayHello(ws *sync.WaitGroup) {
// // 	defer ws.Done()
// // 	a, err:= rand.Int(rand.Reader, big.NewInt(10));
// // 	if err != nil {
// // 		fmt.Println("Error", err)
// // 	}
// // 	fmt.Println("Hello", a.String())
// // }

// func main () {

// 	// var wg sync.WaitGroup
// 	// wg.Add(1)

// 	// defer fmt.Println("Done")
// 	// go SayHello(&wg)
// 	// wg.Wait()

// 	// fmt.Println("World")

// 	// channels
// 	for i := 0; i < 5; i++ {
// 		go func (i int)  {
// 			fmt.Println("Sending::: ", i)
// 		}(i)
// 		time.Sleep(time.Second * 1)
// 	}

// 	for i := 0; i < 5; i++ {
// 		go fmt.Println("Message::: ")
// 	}

// 	time.Sleep(time.Second * 1)
// 	defer fmt.Println("Done")

// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
//     ch := make(chan string, 3)
// 	defer close(ch)
//     go func() {
// 		time.Sleep(time.Second * 1)
//         ch <- "Hello from Goroutine!" // Gửi dữ liệu vào channel

//     }()
// 	go func() {
// 		time.Sleep(time.Second * 1)
// 		ch <- "Hello from Goroutine 2!" // Gửi dữ liệu vào channel

// 	}()
// 	go func() {
// 		time.Sleep(time.Second * 1)
// 		ch <- "Hello from Goroutine 3!" // Gửi dữ liệu vào channel

// 	}()

//     msg := <-ch // Nhận dữ liệu từ channel
//     msg1 := <-ch // Nhận dữ liệu từ channel
// 	msg2 := <-ch // Nhận dữ liệu từ channel
//     fmt.Println(msg)
// 	fmt.Println(msg1)
// 	fmt.Println(msg2)

// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- "Data từ channel 1"
// 	}()
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch2 <- "Data từ channel 2"
// 	}()

// 	select {
// 	case msg1 := <-ch1:
// 		fmt.Println(msg1)
// 	case msg2 := <-ch2:
// 		fmt.Println(msg2)
// 	}
// }

package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string, msg string, delay time.Duration) {
	time.Sleep(delay)
	ch <- msg
}

func main() {
	// Tạo hai channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Gửi tin nhắn từ hai Goroutine khác nhau
	go sendMessage(ch1, "Message from ch1", 2*time.Second)
	go sendMessage(ch2, "Message from ch2", 1*time.Second)

	// Xử lý tin nhắn từ cả hai channel
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	}

	// Đợi một chút để không kết thúc quá nhanh
	time.Sleep(4 * time.Second)
}

