// package main

// import (
// 	"fmt"
// 	"time"
// )

// func producer(id int, ch chan<- string) {
// 	for i := 0; i < 3; i++ {
// 		ch <- fmt.Sprintf("Producer %d: data %d", id, i)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func main() {
// 	ch := make(chan string)

// 	go producer(1, ch)
// 	go producer(2, ch)

// 	for i := 0; i < 6; i++ {
// 		fmt.Println(<-ch) // Nhận dữ liệu từ nhiều producer
// 	}

//		time.Sleep(time.Second * 3)
//	}
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Tạo 3 worker
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Gửi 10 công việc vào channel
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs) // Đóng channel sau khi gửi xong

	wg.Wait() // Đợi tất cả worker hoàn thành
}