package utils

import "fmt"

type FamilyAccount struct {
	//用户输入变量
	key string
	//循环控制器
	loop bool
	//账户余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//是否有收支行为
	flag bool
	//收支详情
	details string
}

// 工厂模式的构造方法
func NewFailyAccount() *FamilyAccount {
	return &FamilyAccount{
		//用户输入变量
		key: "",
		//循环控制器
		loop: true,
		//账户余额
		balance: 10000.0,
		//每次收支的金额
		money: 0.0,
		//每次收支的说明
		note: "",
		//是否有收支行为
		flag: false,
		//收支详情
		details: "收支\t账户余额\t收支金额\t说     明",
	}
}

// 显示明细
func (this *FamilyAccount) showDetails() {
	fmt.Println("---------------当前收支明细记录-----------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧！")
	}
}

// 登记收入
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//	收入说明拼接
	this.details += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 登记支出
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	//必须判断
	if this.money > this.balance {
		fmt.Println("余额的金额不足")
		//break
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	//	收入说明拼接
	this.details += fmt.Sprintf("\n支出\t\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 退出系统
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你输入的有误，请重新输入y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

//结构体方法

// 显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n---------------家庭收支记账软件-----------------")
		fmt.Println("               1.收支明细")
		fmt.Println("               2.登记收入")
		fmt.Println("               3.登记支出")
		fmt.Println("               4.退出软件")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项。。。")
		}
		if !this.loop {
			break
		}
	}
}
