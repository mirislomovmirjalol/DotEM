package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"
)

var pushCmd = &cobra.Command{
	Use:     "push",
	Short:   "Push environment to .EM",
	Long:    `Push command designed to push environment from local environment file to .EM`,
	Run:     handlePush,
	Example: ".em push <project_name> <file_path>",
	Args:    cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func handlePush(_ *cobra.Command, args []string) {
	projectName := args[0]
	filePath := args[1]
	err := service.UpdateProject(projectName, filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data pushed successfully")
}
