package main

import "fmt"

type Human struct {
	name interface{}
	word interface{}
}

func (h Human) sayName() string {
	return fmt.Sprintf("%v", h.name)
}

func (h Human) saySmth() string {
	return fmt.Sprintf("%v", h.word)
}

type Action struct {
	Human
	item string
}

func (a Action) act() {
	fmt.Printf("%s acting with a %s and saying: '%s'\n", a.sayName(), a.item, a.saySmth())
}

func main() {
	a := Action{Human: Human{"John", "Ogogo"}, item: "chair"}

	a.act()
}
