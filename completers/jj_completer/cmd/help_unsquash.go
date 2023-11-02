package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_unsquashCmd = &cobra.Command{
	Use:   "unsquash",
	Short: "Move changes from a revision's parent into the revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_unsquashCmd).Standalone()

	helpCmd.AddCommand(help_unsquashCmd)
}
