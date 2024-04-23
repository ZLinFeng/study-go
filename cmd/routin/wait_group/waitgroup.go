package waitgroup

import "sync"

func run(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		println("Goroutin ", num, ": print ", i)
	}
}

func BasicUse() {
	// 类似于Java里面的 CountDownLaunch
	var wg sync.WaitGroup

	// 启动5个协程
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go run(i, &wg)
	}
	wg.Wait()
	println("All finished.")
}
