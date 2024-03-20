package main
import (
	"fmt"
)

type Animal interface {
	Name() string
	Speak() string
	Play()
}

// class Dog
type Dog struct {
	name string
	gender string
}

func (d *Dog) Play(){
	fmt.Println(d.Speak())
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("my name is %v and my gender is %v", d.name, d.gender)
}

func (d *Dog) Name() string {
	return d.name
}

// class Cat
type Cat struct {
	name string
	gender string
}

func (d *Cat) Play(){
	fmt.Println(d.Speak())
}

func (d *Cat) Speak() string {
	return fmt.Sprintf("I'm a cat, my name is %v and my gender is %v", d.name, d.gender)
}

func (d *Cat) Name() string {
	return d.name
}

func Play(a Animal) {
	a.Play()
}

func main(){
	d :=&Dog{"KeJi", "Male"}
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
	Play(d)

	c :=&Cat{"LittleC", "Male"}
	fmt.Println(c.Name())
	fmt.Println(c.Speak())
	Play(c)
}