package main

import "fmt"
func main(){
	nums:= [4]int{2,7,11,15}
	var target int = 9
	res:=twoSum(nums[:],target)
	fmt.Println(res)


}
func twoSum(nums []int, target int) []int{
    var length int = len(nums)
    //fmt.Println("len",length)
    var my_Map= make(map[int]int)
    for i:=0; i<length;i++{
        my_Map[nums[i]]=i
    }
    //fmt.Println(my_Map)
    
    for i:=0;i<length;i++{
        var diff int =target -nums[i]
        val,key:=my_Map[diff]
            //fmt.Println(key,val)

        if key && i!= val {
            return []int{i, val}
            
        }
        
    }
    
    return []int{}
    
}