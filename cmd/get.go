package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"
)

var getCmd = &cobra.Command{
	Use:  "get",
	Args: cobra.ExactArgs(1),
	Run:  getProject,
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getProject(cmd *cobra.Command, args []string) {
	projectName := args[0]

	data, err := service.GetProject(projectName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
