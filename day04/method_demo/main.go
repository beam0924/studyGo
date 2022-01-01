package main

import "fmt"

type dog struct {
	name string
}

type person struct {
	age int
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func newPerson(age int) person {
	return person{
		age: age,
	}
}

//方法是作用于特定类型的函数
//接收者表示的是接收方法的具体类型变量，它多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~\n", d.name)
}

func (p person) guonian() {
	p.age++
}

func (p *person) zhenguonian() {
	p.age++
}

func main() {
	d1 := newDog("beam")
	d1.wang()
	p1 := newPerson(17)
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	fmt.Println("-------------")
	fmt.Println(p1.age)
	p1.zhenguonian()
	fmt.Println(p1.age)
}
