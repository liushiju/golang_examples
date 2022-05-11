package studentinfo

import "fmt"

// 学生有id、姓名、年龄、分数
// 定义student类型
type Student struct {
	id    int64
	name  string
	age   uint
	score uint
}

// 学生管理
type MgmStu struct {
	allStu []*Student
}

// 构造出学生信息函数
func NewStu(id int64, name string, age, score uint) *student {
	return &Student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}

// 构造出学生管理函数
func NewMgmStu() *MgmStu {
	return &MgmStu{
		allStu: make([]*Student, 0, 100),
	}
}

// 添加学生
func (m *MgmStu) AddStu(newStu *Student) {
	m.allStu = append(m.allStu, newStu)
}

// 编辑学生
func (m *MgmStu) EditStu(newStu *Student) {
	for i, v := range m.allStu {
		if newStu.id == v.id {
			m.allStu[i] = newStu
			return
		}
	}
	fmt.Printf("系统中没有学号为%v的学生\n", newStu.id)
}

// 删除学生
func (m *MgmStu) DelStu(id int64) {
	for i, v := range m.allStu {
		if id == v.id {
			// 删除切片必须这么删
			m.allStu = append(m.allStu[:i], m.allStu[i+1:]...)
		}
	}
	fmt.Printf("系统中没有学号为%v的学生\n", id)
}

// 显示全部学生
func (m *MgmStu) ShowStu() {
	for _, v := range m.allStu {
		fmt.Printf("学号: %d, 姓名: %s, 年龄: %d, 分数:%d\n", v.id, v.name, v.age, v.score)
	}
}
