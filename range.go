package main
import "fmt"
func main() {
	entries := []int{10,20,30,40,50}
	for _,val := range entries{
		fmt.Println(val)
	}
	//why?
	// for val := range entries{
	// 	fmt.Println(val)
	// }
	scores := map[string]int{
		"wade":0,
		"david":1,
		"DW":3,
	}
	for name, score := range scores{
		fmt.Println(name , ":", score)
	}
}