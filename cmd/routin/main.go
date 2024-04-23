package main

import (
	"time"

	"github.com/ZLinFeng/study-go/cmd/routin/channels"
	selectuse "github.com/ZLinFeng/study-go/cmd/routin/select"
	waitgroup "github.com/ZLinFeng/study-go/cmd/routin/wait_group"
)

func routin() {
	for i := range 2 {
		println(i, " in rountin")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go routin()
	for i := 0; i < 2; i++ {
		println(i, " in main")
		time.Sleep(time.Second * 1)
	}

	waitgroup.BasicUse()
	channels.BasicUse()
	selectuse.BasicUse()

	channels.BroadCast()
}
