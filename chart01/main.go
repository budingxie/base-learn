package main

import (
	"base-learn/chart01/array"
	"base-learn/chart01/pymap"
)

func main() {
	// 1.数组申明
	array.State()

	// 2.数组使用
	array.UseArray()

	// 3.数组复制
	array.UseStrArr()

	// 4.错误复制
	array.AssignArr()

	// 5.指针数组复制
	array.UsePoint()

	// 6.二维数组
	array.MultiDimArr()

	// 7.二维数组访问
	array.UseMultiDimArr()

	// 8.切片的创建
	array.SliceCreate()

	// 9.使用切片
	array.UseSlice()

	// 10.遍历切片
	array.IterationSlice()

	// 11.map创建
	pymap.CreateMap()

	// 12.使用map
	pymap.UseMap()
}
