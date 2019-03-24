package lessons

import (
	"fmt"
	"time"
)

func goroutineTest() {
	//var wg sync.WaitGroup
	//defer wg.Wait()
	//wg.Add(1)
	//
	//go goroutine("world", &wg)
	//normal("Hello")
	//
	//wg.Wait()
}
//func normal(s string) {
//	for i := 0; i<5; i++ {
//		fmt.Println(s)
//	}
//}
//func goroutine(s string, wg *sync.WaitGroup) {
//	for i := 0; i<5; i++ {
//		time.Sleep(100 * time.Millisecond)
//		fmt.Println(s)
//	}
//	wg.Done()
//}


func channelTest() {

	c := make(chan int)

	go goroutine([]int{1,2,3,4,5}, c)
	go goroutine([]int{2,3,4,5,6}, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Print(y)
}

func goroutine(s []int, c chan int) {
	sum := 0
	for _, y := range s {
		sum += y
		time.Sleep(100 * time.Millisecond)
	}
	c <- sum

}