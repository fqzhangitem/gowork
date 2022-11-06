package main

import "fmt"

func main() {
	// 声明变量 var 变量名 类型
	/**
	类型 int int8 int16 int32 int64 float32 float64
	单精度最多有7位有效数字
	*/
	num()
	str()
	copyValue()
}

func copyValue() {
	var x = 10
	// 值拷贝，对应的是Java的深拷贝
	var y = x
	x = 20
	var value = x + y
	fmt.Println(x, y, value)
}

func str() {
	name, age, isEmpty := "张三", 20, true
	fmt.Println(age, name, isEmpty)
	var str = "三"
	fmt.Println(name + str)
}

func num() {
	var (
		a int
		b int8
		c int16
		d int32
		e int64
		f float32
		h float64
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(h)
}
