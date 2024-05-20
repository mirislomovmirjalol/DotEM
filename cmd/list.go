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
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: listProjects,
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
