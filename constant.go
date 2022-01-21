package main
import(
	"fmt"
	"math"
)
const s string ="constant"
func main(){
	const v = 500000;
	fmt.Println(s);
	fmt.Println(math.Sin(v))
}