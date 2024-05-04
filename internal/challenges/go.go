package challenges

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleGo() {
	challenges := []list.Item{
		Item("Challenge 1"),
		Item("Challenge 2"),
	}
	challengesModel := Model{List: CreateList(challenges, "Challenge")}

	if _, err := tea.NewProgram(challengesModel).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	switch ChallengeChosen {
	case "Challenge 1":
		fmt.Println("hey bro you chose: Challenge 1")
	case "Challenge 2":
		fmt.Println("hey bro you chose: Challenge 2")
	}
}
