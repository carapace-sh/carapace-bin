package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "let",
	Short: "Evaluate arithmetic expressions",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-let",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
}
