package BaseAnimal

type BaseAnimal struct {
}

func (b BaseAnimal) MakeNoise() string {
	return "Your animal made noise"
}

func (b BaseAnimal) AnotherFunc() string {
	return "Another func from base animal"
}
