package main

import (
	"fmt"
	"math"
)

func main(){
	num:=20
	res:= isPerfectSquare((num))
	fmt.Println((res))
}
func isPerfectSquare(num int) bool {
    for i:=1;i<=int(math.Pow(2,31))-1;i++{
        if int(math.Pow(float64(i),2)) ==num{
            return true
        }else if i>num{
            return false
		}else{

		}
        
        
    }
    return false
}