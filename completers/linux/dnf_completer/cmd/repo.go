package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo <subcommand> [options] [<repo-spec>...]",
	Short: "manage repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()

	rootCmd.AddCommand(repoCmd)
}
