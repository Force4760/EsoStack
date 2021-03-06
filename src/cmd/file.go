package cmd

import (
	"github.com/force4760/esostack/src/io/interpret"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Run a single file",
	Long:  "Run a single file (provided as an argument) and print the stack at the end",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		interpret.File(args[0])
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
