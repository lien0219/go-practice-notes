package service

import "go_code/customerManage/model"

// 增删改查（操作Customer）
type CustomerService struct {
	customers []model.Customer
	//	当前切片含有多少客户
	//	该字段后面可以作为新客户的id+1
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "20", 50, "18733333333", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添加客户到customers切片
func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

// 根据id查找客户在切片中对应的下标，如果没有返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

// 根据id删除客户（切片中删除）
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	//从切片中删除
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

// 修改
func (this *CustomerService) Update(id int, customer model.Customer) model.Customer {
	index := this.FindById(id)
	if index == -1 {
		return model.Customer{}
	}
	this.customers = append(this.customers, customer)
	return this.customers[index]
}
