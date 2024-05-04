package main

import (
	"fmt"
	"interview/internal/challengelist"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func handleGo() {
	challenges := []list.Item{
		challengelist.Item("Challenge 1"),
		challengelist.Item("Challenge 2"),
	}
	challengesModel := challengelist.Model{List: challengelist.CreateList(challenges, "Challenge")}

	if _, err := tea.NewProgram(challengesModel).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	switch challengelist.ChallengeChosen {
	case "Challenge 1":
		fmt.Println("hey bro you chose: Challenge 1")
	case "Challenge 2":
		fmt.Println("hey bro you chose: Challenge 2")
	}
}
