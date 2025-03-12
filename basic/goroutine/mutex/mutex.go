// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter int = 0
// var murw sync.RWMutex

// func readCounter(){
// 	murw.RLock() // Khóa đọc
// 	defer murw.RUnlock()
// 	fmt.Println("Counter:", counter)
// }

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 5; i++ {
// 		murw.Lock()   // Khóa trước khi cập nhật
// 		counter++
// 		murw.Unlock() // Mở khóa sau khi cập nhật
// 	}
// }

// func main() {

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	for i := 0; i < 10; i++ {
// 		go readCounter()
// 	}
// 	go increment(&wg)
// 	go increment(&wg)

// 	wg.Wait()
// 	fmt.Println("Counter:", counter) // Kết quả không phải lúc nào cũng là 2000!
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var once sync.Once
// var counter int = 0

// func setup() {
// 	counter++
// }

// func worker(id int) {
// 	once.Do(setup) // Đảm bảo setup chỉ chạy 1 lần
// 	fmt.Printf("Worker %d started\n", id)
// }

// func main() {
// 	for i := 1; i <= 3; i++ {
// 		go worker(i)
// 	}

// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Counter:", counter) // Kết quả luôn là 1

// 	fmt.Scanln() // Chờ input để giữ chương trình
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

var ready = false
var cond = sync.NewCond(&sync.Mutex{})

func worker(id int) {
	cond.L.Lock()
	for !ready {
		cond.Wait() // Chờ tín hiệu từ main()
	}
	fmt.Printf("Worker %d started\n", id)
	cond.L.Unlock()
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i)
	}

	time.Sleep(time.Second)
	cond.L.Lock()
	ready = true
	cond.Broadcast() // Gửi tín hiệu cho tất cả Goroutine
	cond.L.Unlock()

	time.Sleep(time.Second)
}
