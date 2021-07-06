package loop

import "fmt"

func Loop() {
	total := 100
	//	普通循环体
	for i := 0; i < total; i++ {
		if i > 50 {
			fmt.Println("exit")
			// break 退出循环
			break
		}
		fmt.Println(i)
	}
	//	无限循环
	//for {
	//	fmt.Println("loading...")
	//}
	//	遍历数组
	arr := [...]string{"chen", "huang", "ming", "sb"}
	//	直接遍历一个数组 通过range可得到k和v  k -> key值 v -> value值
	for k, v := range arr {
		fmt.Println(k)
		fmt.Println(v)
	}
}
