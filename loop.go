package main
import "fmt"
func main() {
	list :=[5] int {1,2,3,4,5}
	
	// count := 4
	for i := 0; i < len(list); i++ {
		fmt.Println(i, list[i]);
	}
}