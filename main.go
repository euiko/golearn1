package main
import "fmt"

type person struct {
	name string
	age uint64
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

//var human person = person {
//	"euiko",
//	10,
//}

const MAX uint64 = 10


func main() {
	work := make(chan uint64, MAX)
	result := make(chan uint64)

	go func(){
		println("Generating Work value")
		for i := uint64(1); i < MAX ; i++{
			if (i % 3) == 0 || (i % 5) == 0{
				work <- i
			}
		}
		println("Work value have been generated")
		close(work)
	}()


	go func() {
		r := uint64(0)
		s_work := fmt.Sprintf("work length %d", len(work))
		println(s_work)
		println("Counting result")
		s := ""
		for i := range work {
			s += fmt.Sprintf("Work count is %d %d \n ", len(work), i)
			r = r + i
		}
		result <- r
		println(s)

		s_work = fmt.Sprintf("work length %d", len(work))
		println(s_work)
		println("Result counted")
	}()

	println(<-result)
	//fmt.Println("Kharis",<-result)
}
