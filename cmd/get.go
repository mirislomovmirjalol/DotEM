package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get project by name",
	Long:  `Get command designed to get project by name from .EM storage`,
	Args:  cobra.ExactArgs(1),
	Run:   getProject,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getProject(_ *cobra.Command, args []string) {
	projectName := args[0]

	data, err := service.GetProject(projectName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
