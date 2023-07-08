package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 定义一个messages的单通道channel,接收int类型的数据，长度为10
	messages := make(chan int, 10)
	// 定义一个done的单通道channel，接收bool类型的数据，长度为0
	done := make(chan bool)

	//延迟关闭messages通道，等到主协程最后一条命令执行完毕
	defer close(messages)

	// 消费者consumer，定义一个子协程B
	go func() {
		// 以i变量循环20次
		for i := 0; i < 20; i++ {
			// 通过select关键字循环遍历通道messages的数据信息，以及判断通过是否为就绪
			select {
			// 当主协程done通道关闭，通知子协程不要继续消费了
			case <-done:
				fmt.Println("child process interrupt...")
				return
			// 如果主协程正常运行，子协程继续读取数据
			default:
				//等待10秒后才允许消费者去通道里读取数据，检查队列满生产者阻塞情况
				time.Sleep(10 * time.Second)
				fmt.Println("get number is:", <-messages)
				// 如果循环次数超过20次，就退出循环并返回
				if i >= 20 {
					fmt.Println("loopfor more than 100 times, loop for exit!")
					return
				}
			}
		}
	}()

	// 生产者producer,定义一个子协程A
	// 以i变量循环20次
	for i := 0; i < 20; i++ {
		// 通过select关键字循环遍历通道messages的数据信息，以及判断通过是否为就绪
		select {
		// 当主协程done通道关闭，通知主协程不要继续生产了
		// 让主协程直接关闭，检查队列空消费者阻塞情况
		case <-done:
			fmt.Println("main process interrupt...")
			return
		// 如果主协程正常运行，且继续存入数据
		default:
			// a值从0到10中随机赋值
			a := rand.Intn(10)
			// 打印每次a的随机值
			fmt.Println("set number is:", a)
			// 等待10秒后才让生产者往通道里存入数据，检查队列空消费者阻塞情况
			//time.Sleep(10 * time.Second)
			// 把随机值a存入数据到messages通道中
			messages <- a
			// 如果循环次数超过20次，就退出循环并返回
			if i >= 20 {
				fmt.Println("loopfor more than 100 times, loop for exit!")
				return
			}
		}
	}
	// 关闭done通道，通知消费者子协程，主协程中done通道关闭
	close(done)
	fmt.Println("mian process exit!")
}
