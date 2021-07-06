package funcs

import (
	"errors"
)

//	函数声明 在golang中 函数是一等公民
//	函数也跟变量一样通过名称开头大小写确认私有公有
// func IsMethod(args type)returnType {

// }

// IsMethod 示例 一个简单的不能再简单的函数
func IsMethod(args int) int {
	return args
}

// ReturnFunc 返回多个参数 很清楚 就是指定返回的类型 没有error? 没事 自己通过errors这个包new一个错误
func ReturnFunc() (string, error) {
	return "test", errors.New("错误返回")
}

// varFunc 将返回值当做变量来使用 也是骚操作 Google的开发人员怎么想到的
func varFunc() (data string, err error) {
	data = "test"
	err = errors.New("test")
	return data, err
}
