package child

import "fmt"

type Valueer interface {
	Printf()
}

type Value3 struct {
	num int
}

func (self *Value3) Printf() {
	fmt.Println("module project3.com/test")
	fmt.Println(self.num + 80)
}
