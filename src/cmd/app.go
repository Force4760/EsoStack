package cmd

import (
	"github.com/force4760/esostack/src/app"
	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Open the graphical version of EsoStack",
	Long:  "Open a gui (graphical user interface) of the EsoStack programming language",
	Run: func(cmd *cobra.Command, args []string) {
		app.RunApp()
	},
}

func init() {
	rootCmd.AddCommand(appCmd)
}
