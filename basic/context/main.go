package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Nếu Context bị hủy
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(time.Second)
		}
	}
}

func main() {

// 	ctx, cancel := context.WithCancel(context.Background()) // Tạo Context có thể hủy
// 	a:= time.Now()
// 	go worker(ctx) // Truyền Context vào worker
// 	fmt.Println("Before:::  ", <-ctx.Done()) // false
// 	time.Sleep(3 * time.Second)
// 	cancel() // Hủy Context sau 3  khi cancel() được gọi, ctx.Done() sẽ trả về giá trị
// 	fmt.Println("Before:::  ", <-ctx.Done()) // true

// 	time.Sleep(time.Second) // Chờ worker dừng
// 	fmt.Println(time.Since(a))
// }

	fmt.Println(time.Now().Add(4 * time.Second))

	ctx, cancel:= context.WithDeadline(context.Background(), time.Now().Add(4 * time.Second)) // Tạo Context với deadline

	go worker(ctx)
	defer cancel()
	time.Sleep(5 * time.Second) // Chờ worker dừng

}
// package main

// import (
// 	"context"
// 	"fmt"
// )

// func main() {
// 	ctx := context.WithValue(context.Background(), "key", "value")

// 	fmt.Println(ctx.Value("key").(string)) // value
// }