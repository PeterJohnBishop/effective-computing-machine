package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create a new Bubble Tea program
	p := tea.NewProgram(initialModel())

	// Start the program and handle any errors
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	// Print a message indicating that the program has finished
	fmt.Println("Program finished.")
}

type Menu struct {
	Items    []string
	Index    int
	Selected map[int]struct{}
}

func initialModel() Menu {
	// Initialize the menu with some items
	return Menu{
		Items:    []string{"Item 1", "Item 2", "Item 3"},
		Index:    0,
		Selected: make(map[int]struct{}),
	}
}

func (m Menu) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.Index > 0 {
				m.Index--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Index < len(m.Items)-1 {
				m.Index++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.Selected[m.Index]
			if ok {
				delete(m.Selected, m.Index)
			} else {
				m.Selected[m.Index] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m Menu) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.Items {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.Index == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.Selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
