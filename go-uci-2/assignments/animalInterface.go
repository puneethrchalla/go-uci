package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

type Bird struct{}

type Snake struct{}

// Cow methods
func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird methods
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake methods
func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var command, name, arg string
	animals := make(map[string]Animal)
	var animal Animal

	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s %s", &command, &name, &arg)
		if command == "newanimal" {
			switch strings.ToLower(arg) {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Printf("Invalid animal '%s' entered\n", arg)
			}
			if animal != nil {
				animals[name] = animal
				fmt.Println("Created it!")
			}
		} else if command == "query" {
			animal, exist := animals[name]
			if !exist {
				fmt.Println("Animal doesn't exist!")
				continue
			}
			switch arg {
			case "speak":
				animal.Speak()
			case "move":
				animal.Move()
			case "eat":
				animal.Eat()
			default:
				fmt.Println("incorrect information requested!")
			}
		} else {
			fmt.Println("incorrect command entered!")
		}
	}
}
