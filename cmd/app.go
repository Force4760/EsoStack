package cmd

import (
	"github.com/force4760/esostack/app"
	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Open the graphical version of Esostack",
	Long:  "Open a gui (graphical user interface) of the Esostack programming language",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunApp()
	},
}

func init() {
	rootCmd.AddCommand(appCmd)
}
