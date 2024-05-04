package challenges

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func HandleJavascript() {
	challenges := []list.Item{
		Item("Arrow Functions"),
		Item("Closures"),
		Item("Promises"),
		Item("Prototypes"),
		Item("This Keyword"),
	}
	challengesModel := Model{List: CreateList(challenges, "Challenge")}

	if _, err := tea.NewProgram(challengesModel).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	switch ChallengeChosen {
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
