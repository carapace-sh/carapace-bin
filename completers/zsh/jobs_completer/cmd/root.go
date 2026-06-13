package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Display the status of jobs",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-jobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("l", "l", false, "list process IDs")
	rootCmd.Flags().BoolS("p", "p", false, "list only process IDs")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionJobSpecs().FilterArgs(),
	)
}
