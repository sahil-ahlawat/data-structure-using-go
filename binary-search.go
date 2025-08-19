package main
import "fmt"


func main() {
  arr := []int{1,2,3,4,5,8}
  mid := 0
  searchelem := binarySearch(arr, 9);
  fmt.Println(searchelem)
}

func breakArray(arr []int, side string)[]int{
  output := []int{}
  if(len(arr) > 1){
     midpoint := len(arr)/2;
    if side == "left" {
      for i :=0; i < midpoint;i++{
        output = append(output, arr[i])
       
      }
    }else{
      for i := midpoint+1; i < len(arr);i++{
        output = append(output, arr[i])
        
      }
    }
  }
  fmt.Println(output, side)
   return output;
}
func binarySearch(arr []int, tofind int) int{
  if len(arr) > 0 {
    midpoint := len(arr)/2;
    if arr[midpoint] == tofind {
      return midpoint;
    }else if arr[midpoint] > tofind {
     return binarySearch(breakArray(arr, "left"), tofind)
    }else{
      return binarySearch(breakArray(arr, "right"), tofind)
    }
  }else{
    return -1
  }
}
