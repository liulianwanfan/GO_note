package main

import "fmt"

func main() {
	/*
	数组array:
	存储一组相同数据结构
	特点:定长

	切片slice:
	同数组类似，也叫做变长数组或者动态数组
	特点:定长
	是一个引用类型的容器，指向了一个底层数组
	make()
	func make(t Type,size...IntegerType)Type
	第一个参数:类型
	slice，map,chan
	第二个参数:长度len
	实际存储元素的数量
	第三个参数:容量cap
	最多能够存储的元素的数量
	 */
	//1.数组
	arr :=[4]int{1,2,3,4}//定长
	fmt.Println(arr)

	//切片
	var s1 []int
	fmt.Print(s1)

	s2 :=[]int{1,2,3,4}//变长
	fmt.Println(s2)
	fmt.Printf("%T,%T\n",arr,s2)//[4]int,[]int

	s3 :=make([]int,3,8)
	fmt.Println(s3)
	fmt.Printf("容量:%d,长度:%d\n",cap(s3),len(s3))
	s3[0]=1
	s3[1]=2
	s3[2]=3
	fmt.Println(s3)
	//fmt.Println(s3[3])//panic: runtime error: index out of range [3] with length 3

	//append()
	s4:=make([]int,0,5)
	fmt.Println(s4)
	s4 = append(s4,1,2)
	fmt.Println(s4)
	s4 = append(s4,3,4,5,6,7)//自动扩容
	fmt.Println(s4)

	s4 = append(s4,s3...)//把s3里面的所有元素添加进去
	fmt.Println(s4)
	
	//遍历切片
	for i:=0;i<len(s4) ; i++ {
		fmt.Println(s4[i])
	}

	for i,v:=range  s4 {
		fmt.Printf("%d——————>%d\n",i,v)
	}
}
