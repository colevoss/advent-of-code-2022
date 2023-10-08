package elf

type Elf struct {
	Calories int
	Id       int
}

func NewElf(id int) *Elf {
	return &Elf{
		Calories: 0,
		Id:       id,
	}
}

func (e *Elf) AddCall(calories int) {
	e.Calories += calories
}
