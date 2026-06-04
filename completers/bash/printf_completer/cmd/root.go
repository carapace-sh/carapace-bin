package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "printf",
	Short: "Format and print data",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-printf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("v", "v", "", "assign the output to variable name rather than display it")
}
