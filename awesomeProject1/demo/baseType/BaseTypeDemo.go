package main

import (
	"fmt"
	"time"
)

func main()  {
	//基本类型
	var b bool = true
	var i int = 1
	var s string = "abc"
	const con string = "difuwei"
	var arr1 =[3]int{1,2,44} //声明并初始化
	var arr2 [5]string  //声明
	arr2[0] = "a"
	arr2[1] = "b"
	for i := 0; i < len(arr2); i++ {
		arr2[i] = "c"
	}
	fmt.Println(arr2)
	fmt.Println(b)
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(arr1)

	//结构体
	type People struct {
		name string
		age  int
	}
	//结构体对象初始化
	xiaoming := People{
		"xiaoming", 18,
	}
	fmt.Println(xiaoming.age,xiaoming.name)
	var xiaowang People
	xiaowang.age = 21
	xiaowang.name = "xiaowang"
	fmt.Println(xiaoming)

	//切片 [] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3
	 s1 :=[]int{1,2,3}
	 // make初始化切片
	 s2 := make([]int,5,6)
	fmt.Println(s1)
	fmt.Println(s2)

	 numbers := []int{2,3,5,6,7}
	 numbers = append(numbers,11,22,33)
	 fmt.Println(numbers)
	 sum :=0
	for num :=range numbers{
		sum+=num
	}
	fmt.Println("numbers-sum",sum)

	//map初始化
	map1 := map[string] string{"name":"xiaoming","sex":"maile","aaaa":"peo"}
	for k, v := range map1 {
		fmt.Println("k:",k,",v:",v)
		if k == "aaaa" {
			//删除元素
			delete(map1,k)
		}
	}
	fmt.Println(map1)
	var nokiaphone Phone
	nokiaphone = new(NokiaPhone)
	nokiaphone.Call()
	iphone := Iphone{}
	iphone.Call()

	//chan1 := make(chan string)
	// go routun
	//go say("world")
	//say("hello")

	chan1 := make(chan string,2)
	chan1 <- "aaa"
	chan1 <- "bbb"
	chanv := <-chan1
	fmt.Println("chan1 v=:",chanv)
	fmt.Println(len(chan1))

}
func say(s string)  {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}
type Phone interface {
	Call()
}
type NokiaPhone struct {
	width string
	height string
}
func (nokiaPhone NokiaPhone) Call() {
	nokiaPhone.width = "10"
	fmt.Println("nokiaPhone calling,width:",nokiaPhone.width)
}

type Iphone struct {

}

func (iphone Iphone) Call()  {
	fmt.Println("iphone calling")
}
