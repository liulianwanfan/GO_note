package main

import "fmt"

//全局变量
var num = 24//声明了一个全局变量，名字是num，数值
func main(){
	var num  = 20//全局变量 定义了num  局部变量也可以使用 相同名字
	var age  =  18//第六行声明了一个变量，名字为age， 数值是18
	//变量的作用域，其实就是作用范围
	fmt.Println(age)
	//学习平台账号，只能登陆千锋教辅平台上有效，在其他地方无效
	//学习平台账号的作用域仅限于千峰教育平台
	//变量的作用域，从声明处开始，到该变量所处的（）的结尾处
	//var age = 18 : age 的作用域在
}
