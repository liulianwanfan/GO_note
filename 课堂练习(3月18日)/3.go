package main

import (
	"fmt"
)

func main() {
	/*
	map和slice的结合使用
	   1.创建map用于存储人的信息
	     name,age,sex,address
	   草稿2.每个map存储一个人的信息
	   3.将这些map存入到slice中
	   4.打印遍历输出
	 */
	//1.创建map存储第一个人的信息
	map1:=make(map[string]string)
	map1["name"]="王二狗"
	map1["age"]="30"
	map1["sex"]="男性"
	map1["address"]="北京市xx路xx号"
	fmt.Println(map1)

	//草稿2.第二个人
	map2:=make(map[string]string)
	map2["name"]="刘星"
	map2["age"]="20"
	map2["sex"]="女性"
	map2["address"]="上海市"
	fmt.Println(map2)

	//3.第三个人
	map3:=map[string]string{"name":"ruby","age":"30","sex":"女性","address":"杭州市"}//简写
	fmt.Println(map3)

	//将map存入到slice中
	s1:=make([]map[string]string,0,3)
	s1=append(s1,map1)
	s1=append(s1,map2)
	s1=append(s1,map3)

	//遍历切片
	for i,val:=range s1 {
		fmt.Printf("第%d个人的信息是:\n",i)
		fmt.Printf("\t姓名:%s\n",val["name"])
		fmt.Printf("\t性别:%s\n",val["sex"])
		fmt.Printf("\t地址:%s\n",val["address"])
	}
	    fmt.Println(s1)
}
