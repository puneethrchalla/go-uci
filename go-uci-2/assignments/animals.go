package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := make(map[string]Animal)

	// populate animals
	animals["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	// input vars
	var animal, action string

	// take input
	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s", &animal, &action)
		v := animals[animal]
		switch action {
		case "eat":
			v.Eat()
		case "move":
			v.Move()
		case "speak":
			v.Speak()
		default:
			fmt.Println("Incorrect action entered!")
		}
	}
}
