package selectuse

import "time"

func producer1(ch chan int) {
	time.Sleep(time.Millisecond * 100)
	ch <- 1
}
func producer2(ch chan int) {
	time.Sleep(time.Millisecond * 1000)
	ch <- 2
}

// select 语句可以等待多个通信操作, 并处理第一个就绪的
func BasicUse() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer1(ch1)
	go producer2(ch2)

	for i := 0; i < 2; i++ {
		select {
		case int1 := <-ch1:
			println("from producer1: ", int1)
		case int2 := <-ch2:
			println("from producer2: ", int2)
		}
	}
}
