package main

import "fmt"

// Problem Statement: Abstract Factory Pattern
// It basically entails where the factory itself could be giving out different items based on something
// In normal use cases this could be an env value, user platform or a user input

type Power interface {
	Attack()
}

type BatmanPower struct{}

func (bp *BatmanPower) Attack() {
	fmt.Println("I am rich, work for me!!!")
}

type SupermanPower struct{}

func (sp *SupermanPower) Attack() {
	fmt.Println("I am just too good, breathing out can kill you!")
}

type Team interface {
	Callout()
}

type BatmanTeam struct{}

func (bt *BatmanTeam) Callout() {
	fmt.Println("Team Orphans")
}

type SupermanTeam struct{}

func (st *SupermanTeam) Callout() {
	fmt.Println("Team immigrants")
}

type Factory interface {
	GetPower() Power
	GetTeamName() Team
}

type BatmanFactory struct{}

func (bf *BatmanFactory) GetPower() Power {
	return &BatmanPower{}
}

func (bf *BatmanFactory) GetTeamName() Team {
	return &BatmanTeam{}
}

type SupermanFactory struct{}

func (sf *SupermanFactory) GetPower() Power {
	return &SupermanPower{}
}

func (sf *SupermanFactory) GetTeamName() Team {
	return &SupermanTeam{}
}

func AbstractFactoryPatternExample(choiceValue string) {
	//	In actual use case, this can come from an env variable or some user input
	var factory Factory
	if choiceValue == "Batman" {
		factory = &BatmanFactory{}
	} else if choiceValue == "Superman" {
		factory = &SupermanFactory{}
	}

	factory.GetPower().Attack()
	factory.GetTeamName().Callout()
}
