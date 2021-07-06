package statement

import "fmt"

//	全局变量
//	go中变量的共享状态是通过首字母大小写确认
//	如为大写说明此变量是一个公有变量 小写开头则为一个私有变量 只在本包能够读取
//	公有变量 所有包都可读取
var (
	// 不设置默认值
	SystemName int // SystemName = 3200 可直接赋值默认值
	//	私有变量
	systemId = 2577 // 不设置默认值 只有本包才可读取
)

// 单个声明
var systemData int
var systemClass string

func Test() {
	// 简短声明 -> var testData int = 1234
	// 根据值的类型会自动推导类型
	testData := 1234
	testStr := "14612132"
	fmt.Print(testData)
	fmt.Println(testStr)
}

// 常量
// 一旦赋值无法改变
const (
	PI = 3.1415926
)

// 声明多个变量
var a, b, c, d int = 1, 2, 3, 4

//	a -> 1 , b -> 2, c -> 3, d -> 4
