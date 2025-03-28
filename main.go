package main

import "fmt"

func getName() string {
	name := ""
	
	fmt.Println("Welcome to Vibhanshu's Casino!")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func main() {
	getName()
}