package main

import (
	"fmt"
	"interview/internal/challengelist"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func handleJavascript() {
	challenges := []list.Item{
		challengelist.Item("Arrow Functions"),
		challengelist.Item("Closures"),
		challengelist.Item("Promises"),
		challengelist.Item("Prototypes"),
		challengelist.Item("This Keyword"),
	}
	challengesModel := challengelist.Model{List: challengelist.CreateList(challenges, "Challenge")}

	if _, err := tea.NewProgram(challengesModel).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	switch challengelist.ChallengeChosen {
	case "Arrow Functions":
		fmt.Println("hey bro you chose: Arrow Functions")

	case "Closures":
		fmt.Println("hey bro you chose: Closures")

	case "Promises":
		fmt.Println("hey bro you chose: Promises")

	case "Prototypes":
		fmt.Println("hey bro you chose: Prototypes")

	case "This Keyword":
		fmt.Println("hey bro you chose: This Keyword")
	}
}
