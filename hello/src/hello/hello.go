package main

import (
	"chapter3"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"oops"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var count int
var lock sync.Mutex
var waitGroup sync.WaitGroup
var c, python, java bool

func main() {
	xiaohui := &Dog{}
	xiaohui.Age = 15
	xiaohui.Name = "xiaohui"
	fmt.Println("xiaohui: ", *xiaohui)
}

type Annimal struct {
	Name string
}

type Dog struct {
	Annimal
	Age int
}

type GxfInt int

func (gxfInt *GxfInt) ChangeGxfIntValue() {
	(*gxfInt)++
}

func TestSayHi() {
	s := &Student{
		Name: "guanxiangfei",
		Age:  18,
	}
	s1 := Student{
		Name: "guanxiangfei",
		Age:  18,
	}
	fmt.Println("address of s = ", s)
	fmt.Println("*s = ", *s)
	fmt.Println("s1 = ", s1)
	s.SayHi("morning~")
	s1.SayHi("evening~")
	fmt.Println("*s = ", *s)
	fmt.Println("s1 = ", s1)
}

func TestPoint() {
	count := 0
	var i *int = &count
	var count1 int = count
	*i = 1
	fmt.Println("i = ", i)
	fmt.Println("count1 = ", count1)
	count1 = 2
	fmt.Println("count = ", count)
}

type Student struct {
	Name string
	Age  int
}

func (s *Student) SayHi(message string) {
	fmt.Println("hello name = ", s.Name, ", age = ", s.Age, " ", message)
	s.Age++
}

func TestSlice() {
	s := make([]int, 5, 10)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	fmt.Println("cap(s) = ", cap(s))
	for i := 1; i < 5; i++ {
		s = append(s, 1, 2, 3)
		fmt.Println("s = ", s)
	}
	fmt.Println("cap(s) = ", cap(s))
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()
}

func TestCap() {
	mySlice := make([]int, 5, 20)

	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
}

func TestDeferStack() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p[0] = ", p[0])
	p[0] = 1
	fmt.Println("p[0] = ", p[0])
	var a [10]int
	a[0] = 1
	fmt.Println("a[0] = ", a[0])

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func TestConvert() {
	const s1 = `hello go`
	b := 13
	s := strconv.Itoa(b)
	fmt.Println("s = ", s)
	c, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		fmt.Println("err = ", err)
	}
	fmt.Printf("%T \n", c)
	os := runtime.GOOS
	fmt.Println("os = ", os)
}

func TestVar() {
	var i int
	fmt.Println(i, c, python, java)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func TestSwap() {
	a := "hello"
	b := "word"
	fmt.Println("a = ", a, ", b = ", b)
	a, b = SwapString(a, b)
	fmt.Println("a = ", a, ", b = ", b)
}

func SwapString(a, b string) (string, string) {
	return b, a
}

func TestHttpClient() {
	resp, error := http.Get("http://www.baidu.com")
	if error != nil {
		fmt.Printf("error %s\n", error)
		return
	}
	io.Copy(os.Stdout, resp.Body)
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
