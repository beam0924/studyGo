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
	return &student{id, name}
}

type studentMgr struct {
	allStudent map[int64]*student
}

func newStudentMgr(a map[int64]*student) *studentMgr {
	return &studentMgr{a}
}

func (s *studentMgr) showAllStudent() {
	for k, v := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
	fmt.Printf("一共有%d名学生\n", len(s.allStudent))
}

func (s *studentMgr) addStudent() {
	var (
		addId   int64
		addName string
	)
	fmt.Print("请输入要添加的学生ID:")
	fmt.Scanln(&addId)
	fmt.Print("请输入要添加的学生姓名:")
	fmt.Scanln(&addName)
	s.allStudent[addId] = newStudent(addId, addName)
	fmt.Printf("添加%d:%s成功！\n", addId, addName)
}

func (s *studentMgr) deleteStudent() {
	fmt.Print("请输入要删除的学生ID:")
	var delId int64
	fmt.Scanln(&delId)
	delName := s.allStudent[delId].name
	delete(s.allStudent, delId)
	fmt.Printf("删除学号为%d的学生%s成功!\n", delId, delName)
}

func (s *studentMgr) editStudent() {
	fmt.Print("请输入要修改的学生ID:")
	var editId int64
	fmt.Scanln(&editId)
	_, ok := s.allStudent[editId]
	if !ok {
		fmt.Printf("学号为%d的学生不存在！\n", editId)
	} else {
		fmt.Printf("对学号为%d的学生进行信息修改！\n", editId)
		var (
			editNewId     int64
			editNewString string
		)
		fmt.Print("请输入新的学号：")
		fmt.Scanln(&editNewId)
		fmt.Print("请输入新的姓名：")
		fmt.Scanln(&editNewString)
		delete(s.allStudent, editId)
		s.allStudent[editNewId] = newStudent(editNewId, editNewString)
		fmt.Printf("成功将学号为%d的学生信息修改为学号为%d的学生信息！\n", editId, editNewId)
	}
}

func main() {
	allStudent := make(map[int64]*student, 50)
	mgr := newStudentMgr(allStudent)
	for {
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`1. 查看所有学生
2. 新增学生
3. 删除学生
4. 修改学生
5. 退出系统	`)
		fmt.Print("请输入你要干什么：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项！\n", choice)
		switch choice {
		case 1:
			mgr.showAllStudent()
		case 2:
			mgr.addStudent()
		case 3:
			mgr.deleteStudent()
		case 4:
			mgr.editStudent()
		case 5:
			fmt.Println("再见~~")
			os.Exit(1)
		default:
			fmt.Println("输入有误！")
		}

	}

}
