package main

import (
	"errors"
	"fmt"
)

type Animal interface {
	Speak()
}

func NewAnimal(animalType string) (Animal, error) {
	switch animalType {
	case "dog":
		return &Dog{}, nil
	case "cat":
		return &Cat{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("wrong animal type: %s", animalType))
	}
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("야옹이 야옹.")
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("멍멍이 멍멍.")
}

type Human struct {
	animal Animal
}

func (h Human) AnimalSpeak() {
	fmt.Println("우리집 동물은 이렇게 울어요.")
	h.animal.Speak()
}

func NewHuman(a Animal) Human {
	return Human{animal: a}
}

func main() {
	catHuman, _ := InitHuman("cat")
	catHuman.AnimalSpeak()

	dogHuman, _ := InitHuman("dog")
	dogHuman.AnimalSpeak()

	_, err := InitHuman("no")
	if err != nil {
		fmt.Println(err)
		fmt.Println("나만 동물 업서")
	}
}
