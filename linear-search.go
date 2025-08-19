package main
import "fmt"


func main() {
  arr := []int{1,2,3,4,5,8}
  searchelem := linearSearch(arr, 8);
  fmt.Println(searchelem)
}
func linearSearch(arr []int, tofind int) int{
  output := -1;
  if len(arr) > 0 {
    for i:=0; i<len(arr); i++{
      if(arr[i] == tofind){
        return i;
      }
    }
  }
  return output;
}
