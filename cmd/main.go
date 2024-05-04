package main

import (
	"fmt"
	"interview/internal/challenges"
	"interview/internal/languages"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	items := []list.Item{
		languages.Item("Go"),
		languages.Item("Javascript"),
		languages.Item("Typescript"),
	}
	languagesModel := languages.Model{List: languages.CreateList(items, "Language")}

	if _, err := tea.NewProgram(languagesModel).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	switch languages.LanguageChosen {
	case "Javascript":
		challenges.HandleJavascript()
	case "Go":
		challenges.HandleGo()
	}
}
