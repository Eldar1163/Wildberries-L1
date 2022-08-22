package main

import "fmt"

type Action struct {
	actionName    string
	actionSubject string
	actionObject  string
	Human         // Это встраивание, теперь все поля и методы структуры типа Human доступны из Action напрямую
}

type Human struct {
	name  string
	age   int
	hobby []string
}

func (h *Human) greeting() {
	fmt.Printf("Hey, my name is %s, I'm %d\n", h.name, h.age)
}

func (h *Human) getHobby() {
	fmt.Println("I like:")
	for _, hobby := range h.hobby {
		fmt.Println(hobby)
	}
}

func (a *Action) whatIDo() {
	fmt.Printf("I'm %s a %s by my %s", a.actionName, a.actionObject, a.actionSubject)
}

func main() {
	human := Human{
		name: "John",
		age:  25,
		hobby: []string{
			"Swimming",
			"Reading",
			"Listening music",
		},
	}

	action := Action{
		actionName:    "Reading",
		actionObject:  "Book",
		actionSubject: "Eyes",
		Human:         human,
	}

	action.greeting() // Метод Human
	action.getHobby() // Метод Human
	action.whatIDo()  // Метод Action
}
