package main

import (
	"fmt"
)

func main() {
	//for index, value := range collection {
	// 循环体
	//}
	//其中：

	//collection 表示要遍历的数据结构，可以是数组、切片、字符串、Map 或通道等。
	//index 表示当前遍历元素的索引，如果数据结构不支持索引，则 index 可以省略。
	//value 表示当前遍历元素的值。
	//for 是一个循环语句，用于遍历 collection 中的所有元素。

	//遍历数组
	nums := []int{2, 3, 4}
	sum := 0
	for index, num := range nums {
		sum += num
		fmt.Println("index:", index)
	}
	fmt.Println("sum:", sum)

	for _, num := range nums {
		fmt.Println(num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		//%s是一个占位符，表示字符串

		//还有%d表示整数，%f表示浮点数，%t表示布尔值，%p表示指针，%v表示任意类型
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		//%c表示字符,输出单个字符
		fmt.Printf("%d: %c\n", i, c)
		fmt.Println(i, c)
	}

}
