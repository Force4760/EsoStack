package cmd

import (
	"github.com/force4760/esostack/src/io/initProject"
	"github.com/spf13/cobra"
)

// replCmd represents the repl command
// TODO: long description
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new EsoStack project",
	Long:  `Create a new EsoStack project with the following structure`,
	Run: func(cmd *cobra.Command, args []string) {
		initProject.Init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
