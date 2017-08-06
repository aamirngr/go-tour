package main

import (
	"fmt"
	"math"
)
type lo string
func (x lo) Error() string{
  return fmt.Sprintf(string(x))

}
func Sqrt(x float64) (float64, error) {
	if x>=0{return math.Sqrt(x),nil
		   
		   }else{
				return 0, lo("not allowed")}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
