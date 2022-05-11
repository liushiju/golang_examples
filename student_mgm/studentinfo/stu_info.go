package main

import "fmt"

// 学生有id、姓名、年龄、分数
// 定义student类型
type student struct {
	id    int64
	name  string
	age   uint
	score uint
}

// 学生管理
type mgmStu struct {
	allStu []*student
}

// 构造出学生信息函数
func newStu(id int64, name string, age, score uint) *student {
	return &student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}

// 构造出学生管理函数
func newMgmStu() *mgmStu {
	return &mgmStu{
		allStu: make([]*student, 0, 100),
	}
}

// 添加学生
func (m *mgmStu) addStu(newStu *student) {
	m.allStu = append(m.allStu, newStu)
}

// 编辑学生
func (m *mgmStu) editStu(newStu *student) {
	for i, v := range m.allStu {
		if newStu.id == v.id {
			m.allStu[i] = newStu
			return
		}
	}
	fmt.Printf("系统中没有学号为%v的学生\n", newStu.id)
}

// 删除学生
func (m *mgmStu) delStu(id int64) {
	for i, v := range m.allStu {
		if id == v.id {
			// 删除切片必须这么删
			m.allStu = append(m.allStu[:i], m.allStu[i+1:]...)
		}
	}
	fmt.Printf("系统中没有学号为%v的学生\n", id)
}

// 显示全部学生
func (m *mgmStu) showStu() {
	for _, v := range m.allStu {
		fmt.Printf("学号: %d, 姓名: %s, 年龄: %d, 分数:%d\n", v.id, v.name, v.age, v.score)
	}
}
