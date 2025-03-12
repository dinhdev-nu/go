package main

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)


type Logger struct {
	LOG struct {
		Log string `yaml:"log"`
	} `yaml:"LOG"`
}


func main() {
	data, err:= os.ReadFile("test.yaml")
	if err != nil {
		panic(err)
	}
	
	var logger *Logger
	err = yaml.Unmarshal(data, &logger)
	if err != nil {
		panic(err)
	}

	log:= logger.LOG.Log


	wg:= sync.WaitGroup{}
	// append to file
	file, err:= os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i:= 0; i < 10; i++ {
		wg.Add(1)
		go func(file *os.File, log string, wg *sync.WaitGroup)  {
			
			defer wg.Done()
			_, err := file.WriteString(log + "\n")
			if err != nil {
				fmt.Println(err)
				return
			}	

			
		}(file, log, &wg)
	}
	
	wg.Wait()
	
	defer file.Close()
	defer fmt.Println("File created")
}

// os.Stat(filename) => kiểm tra file có tồn tại không

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	data := []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 1, 91} // Dữ liệu ảnh (ví dụ)
// 	err := os.WriteFile("image.png", data, 0644)
// 	if err != nil {
// 		fmt.Println("Lỗi ghi file ảnh:", err)
// 		return
// 	}
// 	fmt.Println("✅ Ghi file ảnh thành công!")
// }


// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	data, err := os.ReadFile("Screenshot 2025-02-08 111338.png")
// 	if err != nil {
// 		fmt.Println("Lỗi đọc file ảnh:", err)
// 		return
// 	}
// 	fmt.Println("📌 Dữ liệu file ảnh:", data[:20]) // In ra 20 byte đầu tiên
// }