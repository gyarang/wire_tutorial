//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitHuman(animalType string) (Human, error) {
	wire.Build(NewHuman, NewAnimal)
	return Human{}, nil
}
