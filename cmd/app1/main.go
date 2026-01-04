package main

import (
	"fmt"

	"github.com/gyy198/go-plugins-pro/corelib/plugins"
	_ "github.com/gyy198/go-plugins-pro/plugins/plugin_a" // 导入插件使 init 执行
	_ "github.com/gyy198/go-plugins-pro/plugins/plugin_b"
)

func test(v *int, ch chan int) int {
	var sum_v int
	for ch_val := range ch {
		sum_v := *v + ch_val
		fmt.Println("val", *v, sum_v)
	}

	return sum_v
}

func test_select(v *int, ch <-chan int, ch2 <-chan int) {
LOOP:
	for {
		var v1, v2 int
		select {

		case v1 = <-ch:
			fmt.Printf("v1 = %d", *v+v1)
			if v1 > 100 {
				break LOOP
			}
		case v2 = <-ch2:
			fmt.Printf("v1 = %d", *v+v2)
			if v2 > 100 {
				break LOOP
			}

		}

	}
}

func main_test() {
	var i, j int = 1, 200
	ch := make(chan int)
	ch2 := make(chan int)
	go test(&i, ch)
	go test(&j, ch2)
	for c := 1; c < 100; c++ {
		ch <- c
		ch2 <- c + 2
	}

	close(ch)
	close(ch2)

}

func main_test_select() {
	var i, j int = 1, 200
	j = 2
	ch := make(chan int)
	ch2 := make(chan int)
	go test_select(&i, ch, ch2)
	go test_select(&j, ch, ch2)
	for c := 1; c < 100; c++ {
		ch <- c
		ch2 <- c + 2
	}

	close(ch)
	close(ch2)

}

func main() {
	fmt.Println("App1 starting...")

	for _, p := range plugins.All() {
		fmt.Println("Running", p.Name())
		p.Run()
	}

	fmt.Println("App1 finished.")
	fmt.Println("App1 running...")
	main_test()
	// 阻塞程序，不让退出
	fmt.Println("Press Enter to exit")
	fmt.Scanln()
}
