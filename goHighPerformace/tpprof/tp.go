package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

// bigBytes allocates 10 sets of 100 megabytes
func bigBytes() *[]byte {
	s := make([]byte, 1000000)
	return &s
}

func main() {
	var wg sync.WaitGroup

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}

	wg.Add(1)
	wg.Wait() // 为了 pprof 分析的正常运行，阻止 `main` 函数退出
}
