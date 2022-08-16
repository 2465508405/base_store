package main

import (
	"fmt"
	"project/allfunc/test_info/pack"
)

type Parent struct {
	Stu
}
type Stu struct {
	Id int
}

func (s *Stu) add() {
	fmt.Println("aasss")
}

var h int

func main() {
	fmt.Println("ceshi")
	UserGroup := pack.UserGroup{}
	UserGroup.SysApi.Id = 1

	fmt.Printf("%+v", UserGroup)

	s := Parent{}
	s.Stu = Stu{Id: 2}

	s.add()
	s.Stu.add()
}
