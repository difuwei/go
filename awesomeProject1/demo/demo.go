package main

import (
	"demo/demopkg"
	//"fuweidi.com/foo"
	//"fuweidi.com/foo2"
)

func main()  {
	//不同项目下引用本地包
	//fmt.Println(foo.Greet("xiaoming"))
	//同一项目下引用包
	//foo2.Foo2Say()
	demopkg.Say()
}


