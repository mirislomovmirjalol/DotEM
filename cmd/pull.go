package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull environment from DotEM",
	Long:  `Pull environment from DotEM`,
	Run:   handlePull,
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
