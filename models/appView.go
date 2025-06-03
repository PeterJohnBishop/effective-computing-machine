package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppView int

type AppModel struct {
	currentView  AppView
	login        Login
	mainMenu     MainMenu
	AWSMenu      AWSMenu
	ClickUpMenu  ClickUpMenu
	PostgresMenu PostgresMenu
	OpenAIMenu   OpenAIMenu
}

const (
	ViewLogin AppView = iota
	ViewMainMenu
	ViewAWSMenu
	ViewClickUpMenu
	ViewPostgresMenu
	ViewOpenAIMenu
)

func InitialAppModel() AppModel {
	return AppModel{
		currentView:  ViewLogin,
		login:        InitialLogin(),
		mainMenu:     MainMenu{},
		AWSMenu:      AWSMenu{},
		ClickUpMenu:  ClickUpMenu{},
		PostgresMenu: PostgresMenu{},
		OpenAIMenu:   OpenAIMenu{},
	}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.login.Init(),
		m.mainMenu.Init(),
		m.AWSMenu.Init(),
		m.ClickUpMenu.Init(),
		m.PostgresMenu.Init(),
		m.OpenAIMenu.Init(),
	)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case LoginSuccessMsg:
		m.mainMenu = InitialMainMenu(msg.Token, msg.RefreshToken, msg.User)
		m.currentView = ViewMainMenu
		return m, nil

	case MainMenuMsg:
		switch msg.selected {
		case 0: // Postgres
			m.mainMenu = InitialPostgresMenu(msg.token, msg.refreshToken, msg.user)
			m.currentView = ViewPostgresMenu
			return m, nil
		case 1: // OpenAI
			m.mainMenu = InitialOpemAIMenu(msg.token, msg.refreshToken, msg.user)
			m.currentView = ViewOpenAIMenu
			return m, nil
		case 2: // AWS
			m.mainMenu = InitialAWSMenu(msg.token, msg.refreshToken, msg.user)
			m.currentView = ViewAWSMenu
			return m, nil
		case 3: // ClickUp
			m.mainMenu = InitialClickUpMenu(msg.token, msg.refreshToken, msg.user)
			m.currentView = ViewClickUpMenu
			return m, nil
		default:
			m.mainMenu = InitialMainMenu(msg.token, msg.refreshToken, msg.user)
			m.currentView = ViewMainMenu
			return m, nil
		}
	}

	switch m.currentView {
	case ViewLogin:
		updatedLogin, cmd := m.login.Update(msg)
		m.login = updatedLogin.(Login)
		return m, cmd

	case ViewMainMenu:
		updatedMenu, cmd := m.mainMenu.Update(msg)
		m.mainMenu = updatedMenu.(MainMenu)
		return m, cmd

	case ViewAWSMenu:
		updatedAWSMenu, cmd := m.AWSMenu.Update(msg)
		m.AWSMenu = updatedAWSMenu.(AWSMenu)
		return m, cmd
	case ViewClickUpMenu:
		updatedClickUpMenu, cmd := m.ClickUpMenu.Update(msg)
		m.ClickUpMenu = updatedClickUpMenu.(ClickUpMenu)
		return m, cmd
	case ViewPostgresMenu:
		updatedPostgresMenu, cmd := m.PostgresMenu.Update(msg)
		m.PostgresMenu = updatedPostgresMenu.(PostgresMenu)
		return m, cmd
	case ViewOpenAIMenu:
		updatedOpenAIMenu, cmd := m.OpenAIMenu.Update(msg)
		m.OpenAIMenu = updatedOpenAIMenu.(OpenAIMenu)
		return m, cmd
	}

	return m, nil
}

func (m AppModel) View() string {
	switch m.currentView {
	case ViewLogin:
		return m.login.View()
	case ViewMainMenu:
		return m.mainMenu.View()
	case ViewAWSMenu:
		return m.AWSMenu.View()
	case ViewClickUpMenu:
		return m.ClickUpMenu.View()
	case ViewPostgresMenu:
		return m.PostgresMenu.View()
	case ViewOpenAIMenu:
		return m.OpenAIMenu.View()
	default:
		return "Unknown view"
	}
}
