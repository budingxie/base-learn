package main

import "base-learn/chart02/cus"

func main() {

	// 1.结构体类型
	cus.TypeCreate()

	// 2.给自定类型添加方法
	man := cus.Man{
		Name: "lisa",
	}
	// 当关键字func [] funcName之间有参数，这个参数被称为接收者。一个函数有接收者，这个函数就被成为方法
	// 3.1打印
	man.Notify()
	// 3.2修改Email
	man.ChangeEmail("lisa@email.com")
	man.Notify()

	// 4.也可以使用指针来调用使用值接收者申明的方法
	bill := &cus.Man{}
	// golang底层优化了 //(*bill).Notify()
	bill.Notify()
}
