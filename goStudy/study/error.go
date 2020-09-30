package main
import(
	"fmt"
	"math"
	"errors"
)
func sqrt(f float64)(float64, error){
	if f<0 {
		return 0,errors.New("square root of negative number")
	}
	return math.Sqrt(f),nil
}

func main() {
	result, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}
	

}