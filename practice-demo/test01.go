package main
import "fmt"
func main() {
	type Person struct {
		Name string
		Age int
		Gender string
		Address string
		Contact string
	}
	persons := []Person{
		{"foo", 42, "male", "here", "1234"},
		{"FOO", 42, "male", "here", "1234"},
		{"bar", 24, "male", "there", "4321"},
	}
	type StatKey struct {
		Age int
		Gender string
	}
	//st:=make(map[string]int)
	//st["test"]++
	//fmt.Println(st)
	stats := make(map[StatKey]int)
	//fmt.Printf("%#v\n", stats)
	for _, person := range persons {
		stats[StatKey{
			person.Age,
			person.Gender,
		}]++
	}
	fmt.Printf("%#v\n", stats)
}