package main

import (
	"fmt"
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
	languagesModel := languagelist.Model{List: languagelist.CreateList(items, "Language")}

	if _, err := tea.NewProgram(languagesModel).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	switch languagelist.LanguageChosen {
	case "Javascript":
		handleJavascript()
	case "Go":
		handleGo()
	}
}
