package main
import "fmt"
func main(){
	var n, i, m, x, j, final int
	fmt.Scan(&n, &m)
	for i = 1; i <= n; i++{
		if i%m != 0{
			x+=i
		}else{
			j+=i
		}
	}
	final = x-j
	fmt.Print(final)
}