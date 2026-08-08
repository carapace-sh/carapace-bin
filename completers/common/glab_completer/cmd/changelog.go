package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changelogCmd = &cobra.Command{
	Use:   "changelog <command> [flags]",
	Short: "Generate changelogs from your project's commit history.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changelogCmd).Standalone()

	rootCmd.AddCommand(changelogCmd)
}
