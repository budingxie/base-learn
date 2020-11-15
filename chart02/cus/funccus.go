package cus

import "fmt"

// user 在程序里定义一个Man类型
type Man struct {
	Name  string
	Email string
}

// notify使用值接收者实现了一个方法
func (m Man) Notify() {
	fmt.Printf("Sending Man Email To %s<%s>\n",
		m.Name,
		m.Email)
}

// changeEmail使用指针接收者实现了一个方法
func (m *Man) ChangeEmail(email string) {
	m.Email = email
}
