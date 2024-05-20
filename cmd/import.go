package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github/mirislomovmirjalol/DotEM/internal/service"
	"log"
)

var importCmd = &cobra.Command{
	Use:     "import",
	Short:   "Import all data from a file to .EM storage.",
	Long:    `Import command is used to import all data from a file to .EM storage.`,
	Example: ".em import /path/to/file.json",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		err := service.ImportData(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Data imported successfully")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
