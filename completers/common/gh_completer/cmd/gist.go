package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gistCmd = &cobra.Command{
	Use:     "gist <command>",
	Short:   "Manage gists",
	GroupID: "core",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gistCmd).Standalone()

	rootCmd.AddCommand(gistCmd)
}
