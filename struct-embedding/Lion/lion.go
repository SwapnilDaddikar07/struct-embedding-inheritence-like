package Lion

import (
	"embedding/struct-embedding/BaseAnimal"
	"fmt"
)

type Lion struct {
	BaseAnimal.BaseAnimal
}

func (l Lion) Attack() {
	fmt.Println("Lion attacked.")
}

func (l Lion) Jump() {
	fmt.Println("Lion jumped.")
}

func (l Lion) AnotherFunc(){
	fmt.Println("Another func from lion!!")
}
