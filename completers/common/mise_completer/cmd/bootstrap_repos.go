package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Manage Git repositories from `[bootstrap.repos]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_reposCmd).Standalone()

	bootstrap_reposCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_reposCmd)
}
