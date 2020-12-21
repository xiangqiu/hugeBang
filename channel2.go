package main

import (
	"fmt"
	"time"
)

func main()  {
	//create a channel
	ch := make(chan int)

	//开启一个并发匿名函数

	go func() {

		for i := 1; i <= 5; i++{

			ch <- i

			time.Sleep(time.Second/10)
		}
	}()

	for data := range ch{
		fmt.Println(data)

		if data == 5{
			break
		}
	}
}
