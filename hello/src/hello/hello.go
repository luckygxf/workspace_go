package main

import (
	"chapter3"
	"encoding/json"
	"fmt"
	"oops"
	"os"
	"sync"
	"time"
)

var count int
var lock sync.Mutex
var waitGroup sync.WaitGroup

func main() {
	array := []int{3, 4, 5, 6}
	PrintIntArray(array)
	TestArray(array)
	PrintIntArray(array)
}

func TestArray(array []int) {
	array[0] = 1
}

func PrintIntArray(array []int) {
	for _, n := range array {
		fmt.Print(n, " ")
	}
	fmt.Println()
}

func TestOsArgs() {
	len := len(os.Args)
	fmt.Println("len = ", len)
	fmt.Println("os.Args[0] = ", os.Args[0])
}

func TestMap() {
	var ch chan int
	if ch == nil {
		fmt.Println("chan is nil")
		return
	} else {
		fmt.Println("chan is not nil")
	}
	ch <- 1
	var personInfo map[string]string = make(map[string]string)
	personInfo["name"] = "guanxianseng"
	fmt.Println("name = ", personInfo["name"])
}

func TestConcurrent() {
	count = 0
	routineSize := 100000
	waitGroup.Add(routineSize)
	chs := make([]chan int, routineSize)
	for i := 0; i < routineSize; i++ {
		chs[i] = make(chan int, 1)
		go AddCount(chs[i])
	}
	// for i := 0; i < routineSize; i++ {
	// 	// <-chs[i]
	// 	tempCount := <-chs[i]
	// 	fmt.Println("tempCount = ", tempCount)
	// }
	// runtime.Gosched()
	waitGroup.Wait()
	fmt.Println("count = ", count)
}

func AddCount(ch chan int) {
	lock.Lock()
	// fmt.Println("enter lock.lock()")
	count++
	// time.Sleep(10 * time.Second)
	// fmt.Println("count in add count: ", count)
	ch <- count
	defer waitGroup.Add(-1)
	defer lock.Unlock()
}

func TestDeadLock() {
	channel := make(chan string, 2)

	fmt.Println("1")
	channel <- "h1"
	fmt.Println("2")
	channel <- "w2"
	fmt.Println("3")
	channel <- "c3" // 执行到这一步，直接报 error
	fmt.Println("...")
	msg1 := <-channel
	fmt.Println(msg1)
}

func TestMake() map[string]string {
	personInfo := make(map[string]string)
	personInfo["name"] = "guanxianseng"
	personInfo["age"] = "18"
	personInfo["city"] = "beijing"
	return personInfo
}

func TestNew() *int {
	var a int
	return &a
}

func TestPrintJson() {
	guanxianseng := &oops.Person{Name: "guanxiangfei", Age: 18, City: "beijing"}
	// fmt.Printf("%+v\r\n", guanxianseng)
	js, err := json.Marshal(guanxianseng)
	if err != nil {
		fmt.Println("err ", err)
		return
	}
	fmt.Println("js = ", js)
	fmt.Println("string(js) = ", string(js))
}

func TestPrintSpinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

func TestRange() {
	personInfo := map[string]string{"name": "guanxiangfei", "age": "18", "city": "beijing"}
	for k, v := range personInfo {
		fmt.Printf("k = %s, v = %s\n", k, v)
	}
}

func TestPrintf() {
	rect := &oops.Rect{X: 0, Y: 0, Width: 12, Height: 13}
	fmt.Println("area = ", rect.CalArea())
}

func TestPointer() {
	a := 2
	ip := &a

	fmt.Printf("a address = %x\n", &a)
	fmt.Printf("ip value = %x\n", ip)
	fmt.Printf("*ip = %d\n", *ip)
}

func TestOop() {
	rect := new(oops.Rect)
	rect.Width = 10.1
	rect.Height = 12.1
	area := rect.CalArea()
	fmt.Println("rect's area = ", area)
}

func TestSelect() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
			fmt.Println("case ch <- 0")
		case ch <- 1:
			fmt.Println("case ch <- 1")
		}
		fmt.Println("out of select")
		i := <-ch
		fmt.Println("receive value: ", i)
		time.Sleep(time.Second)
	}
}

func PrintEven() {
	n := 0
	for {
		fmt.Println(n)
		n += 2
		time.Sleep(time.Second)
	}
}

func PrintOdd() {
	n := 1
	for {
		fmt.Println(n)
		n += 2
		time.Sleep(time.Second)
	}
}

func ch3() {
	var a chapter3.Integer = 1
	fmt.Println(a.Less(4))
}

func testSwitch(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	}
	switch a {
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	case 3, 4:
		fmt.Println("a = 3 or a = 4")
	}
}
