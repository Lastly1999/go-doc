package controls

import "fmt"

// golang 流程控制
var i = 3

func controlsFunc() {
	// if-else
	if i < 1 {
		fmt.Println("1")
	} else if i < 2 {
		fmt.Println("2")
	} else if i < 3 {
		fmt.Println("3")
	} else {
		fmt.Println("none")
	}
	//	switch 逻辑控制
	switch i {
	case 1:
		fmt.Println("test1")
	case 2:
		fmt.Println("test2")
	case 3:
		fmt.Println("test3")
	default:
		fmt.Println("none")
	}
}
