package main

import "fmt"

type student struct {
	id    int // 学号时唯一的
	name  string
	class string
}

// newStudent 是student类型的构造函数
func newStudent(id int, name string, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学员管理的类型
type studentManagement struct {
	allStudents []*student
}

// newStudentManagement 是studentManagement的构造函数
func newStudentManagement() *studentManagement {
	return &studentManagement{
		allStudents: make([]*student, 0, 100),
	}
}

// 1. 添加学生
func (s *studentManagement) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2. 编辑学生
func (s *studentManagement) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id { // 当学号相同时，就表示找到了要修改的学生
			s.allStudents[i] = newStu // 根据切片的索引直接把新学生的信息赋值
			return
		}
	}
	// 如果走到这里说明输入的学生没有找到
	fmt.Printf("输入的学生信息有误，系统中没有学号是：%d的学生\n", newStu.id)
}

// 3. 展示学生
func (s *studentManagement) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}
