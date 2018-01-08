package main

// Example of OOP in GoLang

import (
	"fmt"

	"github.com/chemt/go/oop/shapes"
)

func main() {
	rect := shapes.NewRectangle(1, 2, 3, 4)

	fmt.Printf("Calling of virtual method String() from Rectangle class:\n\t %s\n", rect)
	fmt.Printf("Calling of virtual method String() from parent (Shape) class:\n\t %s\n", rect.Parent())
}
