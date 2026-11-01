package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var changelogsCmd = &cobra.Command{
	Use:   "changelogs",
	Short: "Show relevant changelogs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changelogsCmd).Standalone()

	rootCmd.AddCommand(changelogsCmd)
}
