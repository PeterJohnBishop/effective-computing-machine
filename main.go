package main

import (
	"fmt"
	"log"

	"effective-computing-machine/main.go/models"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create a new Bubble Tea program
	p := tea.NewProgram(models.InitialAppModel())

	// Start the program and handle any errors
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	// Print a message indicating that the program has finished
	fmt.Println("Program finished.")
}
