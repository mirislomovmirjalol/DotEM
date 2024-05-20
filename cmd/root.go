package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "DotEM",
	Short: ".EM is developer environment management tool",
	Long:  `.EM is a developer environment management tool that helps you to manage your environment files`,
	Run: func(cmd *cobra.Command, args []string) {
		var style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#6b04fd")).Padding(2, 4, 2, 4).MarginBottom(2)
		fmt.Println(style.Render(" __          __  _               z               _              ______ __  __ \n \\ \\        / / | |                            | |            |  ____|  \\/  |\n  \\ \\  /\\  / /__| | ___ ___  _ __ ___   ___    | |_ ___       | |__  | \\  / |\n   \\ \\/  \\/ / _ \\ |/ __/ _ \\| '_ ` _ \\ / _ \\   | __/ _ \\      |  __| | |\\/| |\n    \\  /\\  /  __/ | (_| (_) | | | | | |  __/   | || (_) |    _| |____| |  | |\n     \\/  \\/ \\___|_|\\___\\___/|_| |_| |_|\\___|    \\__\\___/    (_)______|_|  |_|\n                                                                             \n                                                                             "))

		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
