package main

import "fmt"

type animal struct {
	Name  string
	Color string
	Age   int
}

func (a *animal) Run() {
	fmt.Println(a.Name, "is running.")
}

func (a *animal) Jump() {
	fmt.Println(a.Name, "has jumped.")
}

type Cat struct {
	cat animal
}

func main() {
	var cat1 = Cat{
		cat: animal{
			Name:  "Tom",
			Color: "Gray",
			Age:   6,
		},
	}
	fmt.Println(cat1.cat.Name)
	cat1.cat.Run()
	cat1.cat.Jump()
}
