package main

import "fmt"

type AnimalCategory struct {
	order   string // 目
	species string // 种
}

type Animal struct {
	scientificName string // 学名
	AnimalCategory
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s", ac.order, ac.species)
}

func (ac AnimalCategory) hello() string {
	return "hello " + ac.species
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func (a Animal) String() string {
	return fmt.Sprintf("%s (category: %s)", a.scientificName, a.AnimalCategory)
}

type Cat struct {
	name string
	Animal
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func main() {
	category := AnimalCategory{species: "cat"}
	fmt.Printf("animal category: %s\n", category)

	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Println(animal.Category())
	fmt.Println(animal.hello())
	fmt.Println(animal.species)
	fmt.Println(animal)

	cat := Cat{
		name:   "meow",
		Animal: animal,
	}
	fmt.Println(cat)
}
