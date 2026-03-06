package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func ActionPackageSearch(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		rootCmd := cmd.Root()

		_ = rootCmd.Flag("repo")
		_ = rootCmd.Flag("setopt")
		_ = rootCmd.Flag("installroot")

		return carapace.ActionValues() // TODO: implement dnf5 package search
	})
}
