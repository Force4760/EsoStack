package cmd

import (
	"fmt"

	"github.com/Force4760/pipes/io/repl"
	"github.com/spf13/cobra"
)

// Flags
var (
	isLexerOn bool = false
)

// replCmd represents the repl command
var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Read Evaluate Print Loop",
	Long:  `As a Lisp-like language, Pipes need a REPL, it allows you to have a quicker feedback by interpreting lines as you type`,
	Run: func(cmd *cobra.Command, args []string) {
		if isLexerOn {
			repl.Repl()
		} else {
			fmt.Println("REPL")
		}
	},
}

func init() {
	rootCmd.AddCommand(replCmd)
	replCmd.Flags().BoolVarP(&isLexerOn, "lexer", "l", false, "Use the REPL as an interactive Lexer")
}
