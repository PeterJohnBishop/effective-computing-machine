package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ClickUpMenu struct {
	cursor       int
	choices      []string
	selected     map[int]struct{}
	token        string
	refreshToken string
	user         User
	header       string
}

func InitialClickUpMenu(token string, refreshToken string, user User) MainMenu {
	return MainMenu{
		choices: []string{
			"Audit Logs",
			"Authorization",
			"Attachments",
			"Comments",
			"Custom Task Types",
			"Custom Fields",
			"Docs",
			"Folders",
			"Goals",
			"Guests",
			"Lists",
			"Members",
			"Privacy & Access",
			"Roles",
			"Shared Hierarchy",
			"Spaces",
			"Tags",
			"Tasks",
			"Task Checklists",
			"Task Relationships",
			"Templates",
			"Workspaces",
			"User Groups (Teams)",
			"Time Tracking",
			"Time Tracking (Legacy)",
			"Users",
			"Views",
			"Webhooks",
			"Chat (Experimental)",
			"About",
		},
		cursor:       0,
		selected:     make(map[int]struct{}),
		token:        token,
		refreshToken: refreshToken,
		user:         user,
		header:       "Select an API",
	}
}

func (m ClickUpMenu) Init() tea.Cmd {
	return tea.SetWindowTitle("Available APIs")
}

func (m ClickUpMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = make(map[int]struct{})
			m.selected[m.cursor] = struct{}{}
			// switch m.cursor {
			// case 0:
			// 	m.header = "Postgres API Selected"
			// 	return m, func() tea.Msg {
			// 		return MainMenuMsg{selected: 0}
			// 	}
			// case 1:
			// 	m.header = "OpenAI API Selected"
			// 	return m, func() tea.Msg {
			// 		return MainMenuMsg{selected: 1}
			// 	}
			// case 2:
			// 	m.header = "AWS API Selected"
			// 	return m, func() tea.Msg {
			// 		return MainMenuMsg{selected: 2}
			// 	}
			// case 3:
			// 	m.header = "ClickUp API Selected"
			// 	return m, func() tea.Msg {
			// 		return MainMenuMsg{selected: 3}
			// 	}
			// default:
			// 	m.header = "Unknown API Selected"
			// 	return m, func() tea.Msg {
			// 		return MainMenuMsg{selected: -1}
			// 	}
			// }
		}
	}

	return m, nil
}

func (m ClickUpMenu) View() string {
	s := "\nAvailible ClickUp APIs!\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {

			checked = "x"

		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\n\nPress q to quit.\n"

	return s
}
