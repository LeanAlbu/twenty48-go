package main

import(
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/LeanAlbu/twenty48-go/cmd"
)

func main(){
	p := tea.NewProgram(cmd.InitialModel())
	if _, err := p.Run(); err != nil{
		fmt.Printf("Alas, there's an error: %v", err)
		os.Exit(1)
	}
}
