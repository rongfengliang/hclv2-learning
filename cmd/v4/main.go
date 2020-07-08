package main

import (
	"fmt"
	"log"
)

type myFunc func(string)

type myFunc2 func()

func (f myFunc2) Run() {
	f()
}

func (f myFunc) Run(n string) {
	f(n)
}

type job interface {
	Run(string)
}

func app(n string, cmd job) {
	cmd.Run(n)
}

func app2(n string, f func(n string)) {
	app(n, myFunc(f))
}

type user struct {
	name string
}

func (u *user) Exe(n string) {
	log.Println(n)
}
func (u *user) Run(n string) {
	log.Println("from  params----", n)
}

func main() {
	for i := 1; i <= 2; i++ {
		var info string = fmt.Sprintf("%s----%d", "dalong", i)
		u := &user{name: info}
		app2(info, func(n string) {
			u.Run(info)
		})
		//app(string(i), u)
	}
}
