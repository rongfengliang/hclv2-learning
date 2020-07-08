package main

import "fmt"

func main() {
	one(2, callback)
}

//需要传递函数
func callback(i int) {
	fmt.Println("i am callBack")
	fmt.Println(i)
}

//main中调用的函数
func one(i int, f func(int)) {
	two(i, fun(f))
}

//one()中调用的函数
func two(i int, c Call) {
	c.call(i)
}

//定义的type函数
type fun func(int)

//fun实现的Call接口的call()函数
func (f fun) call(i int) {
	f(i)
}

//接口
type Call interface {
	call(int)
}
