package cus

import "fmt"

// 自定义的类型

// 申明一个结构类型
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

type Duration int64

func TypeCreate() {
	// 1.申明user类型的变量，并初始化零值
	var bill user

	// 2.申明user类型的变量，并初始化所有字段
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println(bill)
	fmt.Println(lisa)

	// 3.申明admin类型的变量
	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Println(fred)

	// 4.给不同类型的变量赋值会产生编译错误
	var dur Duration
	//dur = int64(100) //Cannot use 'int64(100)' (type int64) as type Duration in assignment
	dur = Duration(100)
	fmt.Println(dur)
}
