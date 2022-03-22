package cmd

import (
	"github.com/force4760/esostack/src/io/initProject"
	"github.com/spf13/cobra"
)

// replCmd represents the repl command
// TODO: long description
var funcCmd = &cobra.Command{
	Use:   "func",
	Short: "Create a new EsoStack function",
	Long:  `Create and add to the Function Log file a new EsoStack function`,
	Run: func(cmd *cobra.Command, args []string) {
		initProject.FuncFile()
	},
}

func init() {
	rootCmd.AddCommand(funcCmd)
}
