package cmd

import (
	"fmt"
	"os"

	"github.com/Force4760/pipes/io/repl"
	"github.com/spf13/cobra"
)

// Flags
var (
	isLexerOn  bool = false
	isParserOn bool = false
)

// replCmd represents the repl command
var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Read Evaluate Print Loop",
	Long:  `As a Lisp-like language, Pipes need a REPL, it allows you to have a quicker feedback by interpreting lines as you type`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(repl.MSG)

		if isLexerOn {
			repl.Lexer(os.Stdin)
		} else if isParserOn {
			repl.Parser(os.Stdin)
		} else {
			repl.Repl(os.Stdin)
		}
	},
}

func init() {
	rootCmd.AddCommand(replCmd)
	replCmd.Flags().BoolVarP(&isLexerOn, "lexer", "l", false, "Use the REPL as an interactive Lexer")
	replCmd.Flags().BoolVarP(&isParserOn, "parser", "p", false, "Use the REPL as an interactive Parser")
}
