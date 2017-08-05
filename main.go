package main
import "fmt"

type person struct {
	name string
	age int
}

type Stringer interface {
	String() string
}

func header() string{
	return fmt.Sprintf("%-10s %-10s", "Name", "Age")
}

func (p person) String() string{
	return fmt.Sprintf("%-10s %-10d", p.name, p.age)
}

func (p person) Integer() int{
	return p.age
}

var human person = person {
	"euiko",
	10,
}

func main() {
	fmt.Println(header())
	fmt.Print(human)
}
