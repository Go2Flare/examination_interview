package main

import "fmt"

func f(x interface{}){
	fmt.Println(nil, x , x==nil)
	if x == nil{
		fmt.Println(1)
	}else{
		fmt.Println(0)
	}
}

func main() {
	//var x *int = nil
	//f(x)
	//g()
	h()
}


func h(){
	nums := []int{1,2,3,4,5}
	for v := range nums{
		fmt.Println(v)
	}
}

func g(){
	defer func(){
		recover()
	}()
	defer panic("1")
	defer panic("2")
	defer panic("3")
	panic("4")
}

