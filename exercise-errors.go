package main
import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	z := float64(1)
	s := float64(0)
	if x < 0 {
		return math.NaN(),ErrNegativeSqrt(x)
	}else if x == 0 {
		return 0,nil
	}else {
		for i := 0; i < 11 ; i++ {
			z = z - (z*z - x)/(2*z)
			s = z
		}
		return s,nil
	}
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}
func main() {
	fmt.Println(Sqrt(-2))
}
