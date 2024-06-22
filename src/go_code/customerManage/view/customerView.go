package main

import (
	"fmt"
	"go_code/customerManage/model"
	"go_code/customerManage/service"
)

type customerView struct {
	key             string //客户输入
	loop            bool   //循环控制
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (this *customerView) list() {
	//获取切片中的所有用户
	customers := this.customerService.List()
	fmt.Println("----客户列表----")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n---客户列表完成---\n\n")
}

// 添加客户
func (this *customerView) add() {
	fmt.Println("----添加客户----")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	//	新Customer实例
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("-------添加完成-------")
	} else {
		fmt.Println("-------添加失败-------")
	}

}

// 删除
func (this *customerView) delete() {
	fmt.Println("------删除客户------")
	fmt.Println("请选择待删除客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除（Y/N）：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//	调用customerService的Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("-----------删除成功-----------")
		} else {
			fmt.Println("-----------删除失败，输入的id号不存在-----------")
		}
	}
}

// 修改
func (this *customerView) update() {
	fmt.Println("------修改客户------")
	fmt.Println("请选择待修改客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃修改操作
	}

	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	//	新Customer实例
	this.customerService.Update(id, model.NewCustomer2(name, gender, age, phone, email))

	//	调用customerService的Delete方法
	if this.customerService.Delete(id) {

		fmt.Println("-----------修改成功-----------")
	} else {
		fmt.Println("-----------修改失败，输入的id号不存在-----------")
	}

}

// 退出
func (this *customerView) exit() {
	fmt.Println("确认是否退出（Y/N）：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}

		fmt.Println("你输入的有误，确认是否退出（Y/N）：")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

// 主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("--------客户信息管理软件--------")
		fmt.Println("        1. 添 加 客 户         ")
		fmt.Println("        2. 修 改 客 户         ")
		fmt.Println("        3. 删 除 客 户         ")
		fmt.Println("        4. 客 户 列 表         ")
		fmt.Println("        5. 退      出         ")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			//fmt.Println("添 加 客 户")
			this.add()
		case "2":
			//fmt.Println("修 改 客 户")
			this.update()
		case "3":
			//fmt.Println("删 除 客 户")
			this.delete()
		case "4":
			//fmt.Println("客 户 列 表")
			this.list()
		case "5":
			//this.loop = false
			this.exit()
		default:
			fmt.Println("你输入的有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户管理系统...")
}

func main() {
	customerView := customerView{
		key:  "",
		loop: true,
	}

	//	对customerView结构体的customerService字段初始化
	customerView.customerService = service.NewCustomerService()

	//	显示主菜单
	customerView.mainMenu()
}
