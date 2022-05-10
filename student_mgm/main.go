package main

import (
	"fmt"
	"os"
	// "github.com/liushiju/golang_examples/student_mgm/"
)

func showMenu() {
	fmt.Println("==== 学生信息管理系统 ====")
	fmt.Println("1.显示所有学生信息")
	fmt.Println("2.添加学生信息")
	fmt.Println("3.编辑学生信息")
	fmt.Println("4.删除学生信息")
	fmt.Println("5.退出系统")
}

func getInput() *student {
	var (
		id    int64
		age   uint
		name  string
		score uint
	)

	fmt.Println("请按要求输入学生信息")
	fmt.Print("请输入学生学号: ")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的年龄：")
	fmt.Scanf("%d\n", &age)
	fmt.Print("请输入学员的分数：")
	fmt.Scanf("%d\n", &score)
	stu := newStu(id, name, age, score)
	return stu
}

func main() {
	sm := newMgmStu()
	for {
		showMenu()
		var input int
		fmt.Println("按提示输入需要操作的序号")

		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			sm.showStu()
		case 2:
			stu := getInput()
			sm.addStu(stu)

		case 3:
			stu := getInput()
			sm.editStu(stu)

		case 4:
			var id int64
			fmt.Println("输入需要删除学生的学号")
			fmt.Scanf("%d\n", &id)
			sm.delStu(id)
		case 5:
			os.Exit(0)

		}
	}
}
