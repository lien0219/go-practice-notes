package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

// Pupil和Graduate共有的方法绑定到*Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名：%v 年龄：%v 成绩：%v \n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}
func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

type Pupil struct {
	Student //嵌入Student匿名结构体
}
type Graduate struct {
	Student //嵌入Student匿名结构体
}

// 显示成绩
//
//	func (p *Pupil) ShowInfo() {
//		fmt.Printf("学生名：%v 年龄：%v 成绩：%v \n", p.Name, p.Age, p.Score)
//	}
//
//	func (p *Pupil) SetScore(score int) {
//		p.Score = score
//	}
func (p *Pupil) testing() {
	fmt.Println("小学考试中")
}
func (p *Graduate) testing() {
	fmt.Println("大学考试中")
}

func main() {
	//var pupil = &Pupil{
	//	Name: "tom",
	//	Age:  10,
	//}
	//pupil.testing()
	//pupil.SetScore(90)
	//pupil.ShowInfo()

	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(90)
	pupil.Student.ShowInfo()
	fmt.Println("res:", pupil.Student.GetSum(10, 20))

	graduate := &Graduate{}
	graduate.Student.Name = "mary~"
	graduate.Student.Age = 38
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
	fmt.Println("res:", graduate.Student.GetSum(1, 2))
}
