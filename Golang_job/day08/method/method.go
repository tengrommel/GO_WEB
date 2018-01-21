package main

import "fmt"

type Student struct {
	Name string
	Id int
}

func (s *Student)Update(id int)  {
	s.Id = id
}

func (s *Student)UpdateTwo(id int)  {
	s.Id = id
}

func main() {
	var f func(int)
	s := Student{Name:"teng",Id:1}
	f = s.Update
	f(2)
	fmt.Println(s)

	var f1 func(s *Student, id int)
	f1 = (*Student).UpdateTwo
	f1(&s, 3)
	fmt.Println(s)

	s1 := Student{Name:"jack"}
	f1(&s1, 4)
	fmt.Println(s1)
}
