package cmd

import (
	i "github.com/Force4760/pipes/io/initProject"
	"github.com/spf13/cobra"
)

// Flags
var (
	numOfFuncs int = 26
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		i.Init(numOfFuncs)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().IntVarP(&numOfFuncs, "number", "n", 26, "The number of function files to create [1, 26]")
}
