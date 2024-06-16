package main

import "fmt"

// 定义一个结构体Account
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 方法
// 1.存款
func (account *Account) Deposite(money float64, pwd string) {
	//	密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//	存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金钱不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功~~")
}

// 取款
func (account *Account) WithDraw(money float64, pwd string) {
	//	密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//	存款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金钱不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功~~")
}

// 查询余额
func (account *Account) Query(pwd string) {
	//	密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为:%v 余额:%v \n", account.AccountNo, account.Balance)
}

func main() {

	//	测试
	account := Account{
		AccountNo: "lien666",
		Pwd:       "666666",
		Balance:   100.0,
	}

	var G string
	for {
		fmt.Println("请输入功能")
		fmt.Scanln(&G)
		if G == "查询" {
			account.Query("666666")
		} else if G == "存款" {
			account.Deposite(200, "666666")
		} else if G == "取钱" {
			account.WithDraw(150.0, "666666")
		} else {
			return
		}

		//account.Query("666666")
		//account.Deposite(200, "666666")
		//account.Query("666666")
		//
		//account.WithDraw(150.0, "666666")
		//account.Query("666666")
	}
}
