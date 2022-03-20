package main

import (
	"embedding/struct-embedding/BaseAnimal"
	"embedding/struct-embedding/IAnimal"
	"embedding/struct-embedding/Lion"
	"fmt"
)

//Lion struct does not implement the animal interface explicitly by overriding the MakeNoise method.
//Instead, Lion has a field variable of BaseAnimal which in turn implements the AnimalInterface.
//So lion type now implements the Animal Interface.
//Lions field variable should be a non-named field variable i.e just BaseAnimal for this to work.
//We can somehow override the methods as well , it is not overriding per say , but has that effect.
//The AnotherFunc which is on base animal is overriden by lion.
//This is not overriding because the parent class has different method declaration type as compared to the child , parent returns a string , child does not.
func main() {
	ba := BaseAnimal.BaseAnimal{}
	lion := Lion.Lion{
		BaseAnimal: ba,
	}
	fmt.Println(lion)

	lion.AnotherFunc()

	lion.Jump()
	lion.Attack()

	test(lion)
}

func test(animal IAnimal.Animal) {
	fmt.Println(animal.MakeNoise())
}
