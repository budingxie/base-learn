package array

import "fmt"

func State() {
	//申明数组 var name [length] type

	// 1.申明一个包含5个元素的整形数组
	var a1 [5]int

	// 2.申明一个包含5个元素的整形数组
	// 2.1用具体值初始化每个元素
	var a2 [5]int = [5]int{10, 20, 30, 40, 50}

	// 3.申明一个整形数组
	// 3.1使用具体值初始胡化每个元素
	// 3.2容量初始化值由元素数量决定
	a3 := [...]int{10, 20, 30, 40, 50}

	// 4.申明一个由5个元素的数组
	// 4.1使用具体值初始化1和2位置上的元素
	// 4.2其余位置为类型默认值
	// 4.3结果0,10,20,0,0
	a4 := [5]int{1: 10, 2: 20}

	fmt.Println(a1, a2, a3, a4)
}

func UseArray() {
	// 1.申明长度为5的数组
	// 1.1使用具体值为数组赋值
	a1 := [5]int{10, 20, 30, 40, 50}
	// 1.2修改位置为2的数组元素的值
	a1[2] = 35

	v := 2
	// 2.申明包含5个元素的数组，里面的元素指向整数
	// 2.1初始化数组中每个元素的指针的指向
	a2 := [5]*int{0: new(int), 1: new(int), 2: &v}
	// 2.2为索引0，1处位置的元素赋值
	*a2[0] = 10
	*a2[1] = 20

	/**
	------------------------------
	|  地址0  |  地址1  |  地址2  |
	|  指向   |  指向   |  指向   |
	------------------------------
	    |		  |			|
		|		  |			|
		10		  20		2
	*/

	var xx *[]int
	fmt.Printf("%T，%T", xx, a2)
}

func UseStrArr() {
	// 1.什么第一个包含5个元素的字符串数组
	var a1 [5]string

	// 2.申明第二个包含5个元素的字符串数组
	// 2.1用颜色初始化数组
	a2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// 2.2把a2的值复制到a1
	// 复制之后，两个数组的值完全一样
	a1 = a2
	// 数组变量的类型：包括数组长度、每个元素的类型。只有这两部分都相同的数组，才是类型相同的数组，才能相互复制

	a2[0] = "dd"
	fmt.Printf("%v\n%v", a1, a2)
}

func AssignArr() {
	// 1.申明第一个包含4个元素的字符串数组
	//	var a1 [4]string

	// 2.申明第二个包含5个元素的字符串数组
	// 2.1使用颜色初始化数组
	//	a2 := [5] string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// 3.将a2复制给a1
	//	a1 = a2
	//报错
	// compiler Error:
	// cannot use a2 (type [5]string) as type [4]string in assignment
}

func UsePoint() {
	// 1.申明第一个包含3个元素的指向字符串的指针数组
	var a1 [3]*string

	// 2.申明第二个包含3个元素的指向字符串的指针数组
	// 2.1使用字符串指针初始胡这个数组
	a2 := [3]*string{new(string), new(string), new(string)}

	// 3.使用颜色为每一个元素赋值
	*a2[0] = "Red"
	*a2[1] = "Blue"
	*a2[2] = "Green"

	// 4.将a2复制给a1
	a1 = a2
	fmt.Printf("%v,%v\n", a1, a2)

	// 5.修改任意数组中的值
	*a1[0] = "Yellow"
	fmt.Printf("a1[0]=%v, a2[0]=%v\n", *a1[0], *a2[0])
}

func MultiDimArr() {
	// 1.申明一个二维整型数组，两个维度分别存储4个元素和2个元素
	var a1 [4][2]int

	// 2.使用数组字面量来申明并初始化一个二维数组
	a2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

	// 3.申明并初始化外层数组和内存数组的的单个元素
	a3 := [4][2]int{1: {0: 20}, 3: {1: 41}}

	fmt.Printf("a1: %v\na2:%v\na3: %v\n", a1, a2, a3)
}

func UseMultiDimArr() {
	// 1.申明一个2 × 2的二维整型数组
	var a1 [2][2]int

	// 2.设置每个元素的整型值
	a1[0][0] = 10
	a1[0][1] = 20
	a1[1][0] = 30
	a1[1][1] = 40

	// 3.只要类型一致，就可以将多维数组相互赋值
	var a2 [2][2]int
	a2 = a1

	// 4.使用索引为多维数组赋值
	var a3 [2]int = a1[1]
	var value int = a1[1][0]
	a1[1][0] = 100

	fmt.Printf("a1:%v\na2:%v\na3:%v\nvalue:%v\n", a1, a2, a3, value)

	// 当数组作为 方法的参数，最好使用 *[len]int，使用指针可以防止数组复制
}

func SliceCreate() {
	// 1.创建一个字符串切片，其长度和容量都是5个元素
	var slice1 []string = make([]string, 5)

	// 2.创建一个整形切片，其长度为3给元素，容量为5个元素
	slice2 := make([]int, 3, 5)

	// 3.创建一个整型切片，使其长度大于容量。在编译的时候会报错
	//slice3 := make([]int, 5, 3)
	// len larger than cap in make([]int)

	fmt.Printf("slice1:%v\nslice2:%v\n", slice1, slice2)
}

func UseSlice() {
	// 1.创建字符串切片，其长度和容量都是5个元素
	slice1 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// 2.创建一个整型切片，其长度和容量都是3个元素
	slice2 := []int{10, 120, 30}

	// 3.创建字符串切片，使用1字符串初始化第100个元素
	slice3 := []string{99: "1"}
	fmt.Printf("slice1:%v\nslice2:%v\nslice3:%v\n", slice1, slice2, slice3)

	// 4.nil和空切片
	// 4.1创建nil整型切片
	var slice4 []int

	// 4.2使用make创建空的整型切片
	slice5 := make([]int, 0)

	// 4.3使用切片字面量创建空的整型切片
	slice6 := []int{}

	// 不管使用nil切片还是空切片，对其调用内置函数append、len和cap的效果都一样的
	slice4 = append(slice4, 10)

	fmt.Printf("slice4:%v\nslice5%v\nslice6:%v\n", slice4, slice5, slice6)

	slice7 := []int{10, 20, 30, 40, 50}
	slice7[1] = 25

	fmt.Printf("slice7:%v\n", slice7)

	slice8 := []int{10, 20, 30, 40, 50}
	newSlice := slice8[1:3]
	/*
			slice8 := []int{10,20,30,40,50}
			指针地址		长度5		容量5
			    |
				|
		index:	0	1	2	3	4
		value:	10	20	30	40	50
					|
				指针地址		长度2		容量4
			newSlice :=slice[1:3]
			对于newSlice，长度：3-1；容量：5-1
	*/
	// 切片slice8和newSlice 共用一个底层数组，当一个修改了值，另外的也同时变化了
	newSlice[0] = 25
	//newSlice[3] = 35 // index out of range [3] with length 2
	fmt.Printf("slice8:%v\nnewSlice:%v\n", slice8, newSlice)

	// 5.创建切片，并限制容量
	source := []int{10, 20, 30, 40, 50}
	// 从底层数组 index为2，长度为3-2， 容量为4-2
	slice9 := source[2:3:4]
	fmt.Printf("slice9:%v\nsource:%v\n", slice9, source)

	slice9 = append(slice9, 45)
	fmt.Printf("slice9:%v\nsource:%v\n", slice9, source)
}

func IterationSlice() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	// range创建了每个元素的副本，而不是直接返回对该元素的引用
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}

	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}

	// 通过make创建
	s1 := make([]int, 100)
	s2 := make([]int, 0)
	fmt.Printf("s1:%v\n", s1)
	fmt.Printf("s2:%v\n", s2)

	s2 = append(s2, 10)
	fmt.Printf("s2:%v\n", s2)
}
