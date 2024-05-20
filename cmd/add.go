package cmd

import (
	"github.com/spf13/cobra"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add command used for adding new project to DotEM",
	Args:  cobra.ExactArgs(2),
	Run:   addProject,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addProject(_ *cobra.Command, args []string) {
	projectName := args[0]
	filePath := args[1]

	err := service.AddProject(projectName, filePath)
	if err != nil {
		log.Fatal(err)
	}

}
