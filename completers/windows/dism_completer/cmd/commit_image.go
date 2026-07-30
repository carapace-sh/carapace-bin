package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var CommitImageCmd = &cobra.Command{
	Use:   "Commit-Image",
	Short: "apply changes to a mounted image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(CommitImageCmd).Standalone()
	rootCmd.AddCommand(CommitImageCmd)
}
