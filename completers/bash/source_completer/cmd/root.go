package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "source",
	Short: "Execute commands from a file in the current shell",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("p", "p", "", "search path for filename")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"p": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
