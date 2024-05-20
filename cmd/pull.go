package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"
)

var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "Pull environment from .EM",
	Long:    `Pull command designed to pull environment from .EM to local environment file`,
	Args:    cobra.ExactArgs(2),
	Run:     handlePull,
	Example: ".em pull <project_name> <file_path>",
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

func handlePull(_ *cobra.Command, args []string) {
	projectName := args[0]
	filePath := args[1]
	projectData, err := service.GetProject(projectName)
	if projectData == nil {
		log.Fatal("Project not found")
	}
	if err != nil {
		log.Fatal(err)
	}

	err = service.WriteDataToFile(projectData, filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data pulled successfully")
}
