package channels

import (
	"sync"
	"time"
)

// 通过生产者发送消息
func produce(ch chan int) {
	println("开始发送消息")
	for i := range 10 {
		ch <- i
	}
	println("发送完毕")
	close(ch)
}

// 通过消费者处理消息
func consume(ch chan int) {
	for i := range ch {
		println("消费者收到消息: ", i)
	}
}

func BasicUse() {
	// 通过Channels实现多个goroutin的数据传递和同步
	var ch chan int = make(chan int)
	go produce(ch)
	consume(ch)
	println("All finished.")
}

func woker(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second)
	task := <-ch
	println(task)
}

func DynamicWokerPool() {
	// 创建一个容量固定的Channel来缓存任务
	// 队列阻塞 取决于队列容量、发送速度、处理速度
	tasks := make(chan string, 10)

	var waitgroup sync.WaitGroup

	// 动态创建woker
	for i := 0; i < 5; i++ {
		waitgroup.Add(1)
		go woker(tasks, &waitgroup)
	}

	// 下发任务
	for i := 0; i < 20; i++ {
		tasks <- "task"
	}
	close(tasks)
	waitgroup.Wait()
}

// 广播用于通知所有等待的goroutine重新开始执行
func BroadCast() {
	// 定义广播变量
	cond := sync.NewCond(&sync.Mutex{})

	var wg sync.WaitGroup
	// 启动协程
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			println("Start goroutine ", num)
			cond.L.Lock()
			// 等待广播通知
			cond.Wait()
			cond.L.Unlock()
			println("Goroutine receive signal ", num)
		}(i)
	}

	// 模拟执行其他逻辑
	time.Sleep(time.Second * 2)
	// 广播信号
	cond.Broadcast()

	wg.Wait()
}
