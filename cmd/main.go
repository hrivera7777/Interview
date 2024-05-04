package main

import (
	"fmt"
	"interview/internal/challengelist"
	"interview/internal/languagelist"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	items := []list.Item{
		languagelist.Item("Go"),
		languagelist.Item("Javascript"),
		languagelist.Item("Typescript"),
	}

	const defaultWidth = 20

	l := list.New(items, languagelist.ItemDelegate{}, defaultWidth, languagelist.ListHeight)
	l.Title = "Choose a language :"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = languagelist.TitleStyle
	l.Styles.PaginationStyle = languagelist.PaginationStyle
	l.Styles.HelpStyle = languagelist.HelpStyle

	m := languagelist.Model{List: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	switch languagelist.LanguageChosen {
	case "Go":
		goChallenges := []list.Item{
			challengelist.Item("Challenge 1"),
			challengelist.Item("Challenge 2"),
		}
		nombre := challengelist.Model{List: challengelist.CreateList(goChallenges, "Challenge")}

		if _, err := tea.NewProgram(nombre).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	}
}
