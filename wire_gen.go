// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitHuman(animalType string) (Human, error) {
	animal, err := NewAnimal(animalType)
	if err != nil {
		return Human{}, err
	}
	human := NewHuman(animal)
	return human, nil
}
