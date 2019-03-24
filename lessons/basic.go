package lessons

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	"log"
	"os"
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

	//forTest()
	//working()
	//pointer()

	goroutineTest()
	channelTest()

	return
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

func getOSName() string {
	return "mac"
}

func forTest() {

	defer fmt.Println("is defer")
	defer fmt.Println("is defer2")

	for i := 0; i < 10; i++ {

		if i == 3 {
			break
		}

		fmt.Println(i)
	}

	languages := []string{"go", "swift", "python"}
	for i, v := range languages {
		fmt.Println(i, v)
	}

	switch getOSName() {
	case "mac":
		//fmt.Println("is mac")
	case "windows":
		fmt.Println("hogehoge")
	}

	file, error := os.Open("hogehoge")
	if error != nil {
		log.Fatalln("error", error)
	}
	defer file.Close()
}

func working() {

	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	min := intsets.MaxInt
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	fmt.Println("minimum value is ", min)

	m := map[string]int {
		"apple": 200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi": 90,
	}

	sum := 0
	for _, v := range m {
		sum += v
	}
	fmt.Println("sum is", sum)
}

func pointer() {

	i := 100
	fmt.Println(i)
	fmt.Println(&i)

	var p *int = &i
	fmt.Println(*p)
	nonStruct()
}

type MyInt int

func (i MyInt) Double() MyInt {
	return i * 2
}

func nonStruct() {

	myInt := MyInt(10)
	fmt.Println(myInt)
	switchAssertion("hoge")
}

type Human interface {
	Say()
}

type Person struct {
	Name string
}

func (p Person) Say() {
	fmt.Println(p.Name)
}

func switchAssertion(i interface{}) {
	//fmt.Println("hogehoge")
	//ii := i.(int)
	//ii *= 2
	//fmt.Println(ii)

	fmt.Println("type", "assertion")
	ss := i.(string)
	fmt.Println(ss)

	switch v := i.(type) {
	case string:
		fmt.Println(v)
	case int:
		fmt.Println(v * 2)
	default:

	}

	var person2 = Person2{"Mike", 2}
	fmt.Println(person2)

	error := &UserNotFound{"Mike"}
	fmt.Println(error)
}


type Person2 struct {
	Name string
	age int
}

// Stringer 出力方法を変えられる
func (p Person2) String() string {

	return fmt.Sprintf("My name is %v.", p.Name)
}

type UserNotFound struct {
	Username string
}

// エラーは基本的に参照型で定義する
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found %v", e.Username)
}

