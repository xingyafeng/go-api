package main

import "fmt"

type Person struct {
	Name string
}

// 定义结构体，与C语言类似
type Saiyan struct {
	*Person
	Name   string
	Power  int
	Father *Saiyan
}

// go 语言只存在值传递
func Super(s *Saiyan) {

	s.Power += 1
	// s = &Saiyan{Name: "tct", Power: 1}
}

func (p *Person) IntrodUCE() {

	fmt.Printf("I'am is %s\n", p.Name)
}

func (p *Person) setName(s string) {

	p.Name = s
}

func (p *Person) getName() string {

	return p.Name
}

// 结构体的构造器
func NewSaoyan(name string, power int) Saiyan {

	return Saiyan{
		Name:   name,
		Power:  power,
		Father: &Saiyan{Name: "T", Power: power + 1},
	}
}

func main() {

	goku_ := new(Saiyan)
	goku_.Person = &Person{Name: "123"}
	goku_.Name = "tct"
	goku_.Power = 1

	// new vs &  is the same

	goku := &Saiyan{
		Person: &Person{"HAHA"},
		Name:   "Goku",
		Power:  100,
	}

	goku.setName("setname")
	fmt.Println("@@@ ", goku.getName())

	Super(goku_)
	fmt.Println("person :", goku_.Person.Name)
	fmt.Println("person :", goku.Person.Name)
	goku_.IntrodUCE()
	goku.Person.IntrodUCE()
	fmt.Println("power_ : ", goku_.Power)
	fmt.Println("power  : ", goku.Power)
	fmt.Println(NewSaoyan("tct", 1))
	fmt.Println(NewSaoyan("tct", 1).Father)
	fmt.Println(NewSaoyan("tct", 1).Father.Name)
	fmt.Println(NewSaoyan("tct", 1).Father.Power)
}
