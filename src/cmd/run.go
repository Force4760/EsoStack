package cmd

import (
	"github.com/force4760/esostack/src/io/interpret"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the main file",
	Long:  "Run the main file and print the stack at the end",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		interpret.RunMain("./main")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
