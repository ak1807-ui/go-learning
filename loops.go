package main

import "fmt"

func main(){
	for i:=1;i<=5;i++{
		fmt.Println("Count:", i)
}
	num:= []int{10,20,30}
	for index, values := range num{
		fmt.Println("Index:", index, "Values:", values)
	}
}
