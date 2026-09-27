package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Opens the URL of the package's repository in a browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repoCmd).Standalone()

	helpCmd.AddCommand(help_repoCmd)
}
