package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

//var allStudent map[int64]*student
var allStudent = make(map[int64]*student, 50)

func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
	fmt.Printf("一共%d名学生\n", len(allStudent))
}

func addStudent() {
	var (
		addId   int64
		addName string
	)
	fmt.Print("请输入要添加的学生id：")
	fmt.Scanln(&addId)
	fmt.Print("请输入要添加的学生姓名：")
	fmt.Scanln(&addName)
	allStudent[addId] = newStudent(addId, addName)
	fmt.Printf("添加%d:%s成功！\n", addId, addName)
}

func deleteStudent() {
	fmt.Println("请输入要删除的学生的id：")
	var delId int64
	fmt.Scanln(&delId)
	delname := allStudent[delId].name
	delete(allStudent, delId)
	fmt.Printf("删除学号为%d的学生%s成功！\n", delId, delname)
}

func main() {
	for {
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
1. 查看所有学生
2. 新增学生
3. 删除学生
4. 退出系统	
	`)
		fmt.Print("请输入你要干什么：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项！\n", choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			fmt.Println("再见！")
			os.Exit(1)
		default:
			fmt.Println("输入有误")
		}
	}
}
