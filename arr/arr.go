package arr

import "fmt"

func ArrFunc() {
	//	定义数组
	// [长度]类型 普通定义不需要初始化 此时数组的默认值为0 -> 因为这个数组为int类型
	//var arr [3]int
	//var arr [3]string
	//	简单定义 需要初始化 [1,2,3]
	//arrEasy := [3]int{1, 2, 3}
	//arrEasy := [3]string{"1", "2", "3"}
	//fmt.Println(arr)
	//fmt.Println(arrEasy)
	//	不定义数组长度写法 数组类型的标识内增加... 就可以声明他是一个长度随着内容变换的
	//autoArr := [...]string{"nul", "nuoman", "didi"}
	//fmt.Println(autoArr)

	/**************切片**************/
	//	数组切片 使用make声明一个切片 长度为10
	//arrSplit := make([]int, 10) 定义切片
	arrSplit := make([]int, 2, 4)

	// 第二种定义切片的方法
	//arrSplit := []int{1, 23, 4, 5}

	//	切片的push
	arrSplit = append(arrSplit, 2)

	//	切片常用方法 输出长度
	fmt.Printf("切片长度:%d\n", len(arrSplit))
	//	切片的常用方法 检测剩余容量
	fmt.Printf("剩余长度:%d\n", cap(arrSplit))
	fmt.Println(arrSplit)
}
