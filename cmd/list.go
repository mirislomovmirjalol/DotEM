package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"github/mirislomovmirjalol/DotEM/ui"
	"log"
	"os"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all projects",
	Long:    `List command designed to list all projects from .EM storage`,
	Run:     listProjects,
	Example: ".em list",
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listProjects(_ *cobra.Command, _ []string) {
	projects, err := service.GetAllProjects()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(ui.NewModel(projects), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
