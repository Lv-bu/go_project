package main //把test.go文件归属到main,表示test01.go文件所在包是main，在go中，每个文件都必须归属于一个包

import "fmt" //引入包fmt
func main() { //main是主函数
	//输出内容：test
	fmt.Println("hello, My name is Eastmount!")
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
