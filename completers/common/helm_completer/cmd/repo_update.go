package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update information of available charts locally from chart repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_updateCmd).Standalone()
	repoCmd.AddCommand(repo_updateCmd)
}
