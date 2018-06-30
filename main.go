package main

import "fmt"

func main()  {
	fmt.Println("Hello, world!")

	// 插入
	Put("hello", "world")

	// 读取
	fmt.Println(Get("hello"))

	// 删除
	Delete("hello")

	// 读取
	fmt.Println(Get("hello"))

	defer Save()
}

