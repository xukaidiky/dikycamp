package main

import (
	"fmt"
)

func main() {
	//定义一个数组和两个临时变量
	myArray := [5]string{"i", "am", "stupid", "and", "weak"}
	var temp1, temp2 string = "smart", "strong"
	//用for循环遍历字符串数组["I","am","stupid","and","weak"]
	for index, _ := range myArray {
		fmt.Printf(myArray[index] + " ")
	}
	//换个行
	fmt.Printf("\n")
	//将遍历的字符串数组修改为["I","am","smart","and","strong"]
	for index, _ := range myArray {
		if index == 2 {
			myArray[index] = temp1
		}
		if index == 4 {
			myArray[index] = temp2
		}
		fmt.Printf(myArray[index] + " ")
	}
	//换个行
	fmt.Printf("\n")
	//第二种
	for index, value := range myArray {
		switch value {
		case "stupid":
			myArray[index] = "smart"
		case "weak":
			myArray[index] = "strong"
		}
		fmt.Printf(myArray[index] + " ")
	}
}
