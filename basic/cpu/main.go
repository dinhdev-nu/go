package main

import (
	"log"
	"os"
	"runtime/trace"
)

func heavyTask() {
	for i := 0; i < 1e6; i++ {
		_ = i * i
	}
}

func main() {
	// Mở file lưu trace
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Bắt đầu trace
	trace.Start(f)
	defer trace.Stop()

	// Chạy task nặng
	heavyTask()
}
// run tracing
// go run main.go
// go tool trace trace.out


////////////////////////////////////////

// package main

// import (
// 	"log"
// 	"net/http"
// 	_ "net/http/pprof"
// )

// // profiling là một kỹ thuật để đo lường hiệu suất của chương trình

// func main() {
// 	log.Println("Server chạy trên http://localhost:6060/debug/pprof/")

// 	log.Fatal(http.ListenAndServe("localhost:6060", nil))
// }