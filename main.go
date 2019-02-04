package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"
)

const Pi = 3.14

func init() {
	fmt.Println("Init!")
}

func main() {
	bazz()
	mapLesson()
	fmt.Println(function(10, 20))
	fmt.Println(function2(10, 20))
	fmt.Println(function3(10, 20))
	fmt.Println("Hello World?", time.Now())
	fmt.Println(user.Current())

	var i int = 0
	var t, f bool = true, false
	var (
		hoge int
		foo int
		bar int = 2
	)

	float := 1.11

	fmt.Println(i, t, f, hoge, foo, bar, float)
	fmt.Printf("%T", float) // type check
	fmt.Printf("\n%v %T\n", Pi, Pi)

	hohoho := func(x int) {
		fmt.Println("closure", x)
	}

	hohoho(100)

	s := "Hello World"
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(`foo
hoge
hoge
hoge
`)

	// Int to float64
	var x int = 1
	xx:= float64(x)
	fmt.Printf("%T %v\n", xx, xx)

	// ASCII to Integer
	s2 := "14"
	i2, _ := strconv.Atoi(s2)
	fmt.Printf("%T %v\n", i2, i2)

	var a [2]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	var b []int = []int{100, 200, 300, 400, 500, 600}
	fmt.Println(b)

	fmt.Println(b[:2])

	work()
}

func bazz() {
	fmt.Println("bazz")

	//c := make([]int, 5) // lenもcapも5で初期化される
	c := make([]int, 0, 5) // lenは0, capは5
	for i := 0; i<5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	b := make([]int, 0, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
}

func mapLesson() {

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["banana"])
	m["grape"] = 300
	fmt.Println(m)
	fmt.Println(m["nothing"])

	v, ok := m["nothing"]
	fmt.Println(v, ok)

}

func function(x int, y int) int {
	return x + y
}

func function2(x , y int) (int, int) {

	return x + y, x - y
}

func function3(x, y int) (result int) {
	result = x * y
	return
}

func work() {
	fmt.Println("WORK!")

	f := 1.11
	fmt.Println(int(f))

	fmt.Println("[5, 6]")

	var m= make(map[string]int)
	m["Mike"] = 20
	m["Nancy"] = 24
	m["Messi"] = 30
	fmt.Printf("%T %v", m, m)
	}