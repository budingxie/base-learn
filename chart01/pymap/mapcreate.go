package pymap

import "fmt"

func CreateMap() {

	// 1.创建一个映射，键的类型是string，值的类型是int
	var dict map[string]int = make(map[string]int)

	// 2.创建一个映射，键和值的类型都是string；使用两个键值对初始化映射
	map2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	fmt.Printf("dict:%v\nmap2:%v\n", dict, map2)

	map3 := make(map[string]int)
	fmt.Printf("map3:%v\n", map3)
}

func UseMap() {
	// 1.创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{}

	// 2.将Red的代码加入到映射
	colors["Red"] = "#da1337"

	// 3.通过申明映射创建一个nil映射
	var map1 map[string]string
	//map1["Rad"] = "tt"
	// 报错：panic: assignment to entry in nil map
	if map1 == nil {
		fmt.Println("map1 is nil")
	}

	// 4.获取键Blue对应的值，存在
	value, exist := colors["Blue"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Printf("[%s] is not exist. value:--%s--\n", "Blue", value)
	}

	// 5.创建一个映射，进行迭代
	col := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, value := range col {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
	fmt.Println()
	// 6.删除键为Coral的值
	delete(col, "Coral")
	for key, value := range col {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}

	// 7.在函数之间传递映射，并不会制造该映射的一个副本。
}
