package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changelogCmd = &cobra.Command{
	Use:     "changelog",
	GroupID: "configuration",
	Short:   "Show changelog and release notes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changelogCmd).Standalone()
	rootCmd.AddCommand(changelogCmd)
}
