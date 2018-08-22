package main

import "fmt"

func main() {
	yukun := user{"a", "b"}
	yukun.say()
	yukun.say1()

}

type user struct {
	name string
	sex  string
}

func (u user) say1() {
	fmt.Println(u)
}
func (u user) say() {
	u.name = "nihao";
	fmt.Println(u.name, u.sex)
}
