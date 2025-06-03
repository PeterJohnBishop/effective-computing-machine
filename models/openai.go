package models

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type OpenAI struct {
	prompt   string
	response string
}

type AskAI struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	inoutLabel string
	inputFunc  InputFunc
	response   string
}

func InitialAskAI(label string, f InputFunc) IdInput {
	m := IdInput{
		inputs: make([]textinput.Model, 1),
	}

	t := textinput.New()
	t.Cursor.Style = cursorStyle
	t.CharLimit = 64
	t.Width = 20
	t.Placeholder = label
	t.Focus()
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle
	t.CharLimit = 64

	m.inputs[0] = t

	m.inputFunc = f

	return m
}

func (m AskAI) Init() tea.Cmd {
	return textinput.Blink
}

func (m AskAI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				val := m.inputs[0].Value()
				if val == "" {
					fmt.Println("Input cannot be empty.")
					return m, nil
				}
				resp, err := m.inputFunc(val)
				if err != nil {
					fmt.Println("Error:", err)
					return m, nil
				}
				return m, func() tea.Msg {
					return resp
				}
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInput(msg)

	return m, cmd
}

func (m *AskAI) updateInput(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m AskAI) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
