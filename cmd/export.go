package cmd

import (
	"fmt"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:     "export",
	Short:   "Export all data from .EM storage to a single json file.",
	Long:    `Export command is used to export all data from .EM storage to a single json file.`,
	Example: `.em export /path/to/file.json`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		err := service.ExportData(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Data exported successfully")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
