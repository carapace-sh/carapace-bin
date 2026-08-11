package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_resolveCmd).Standalone()

	revision_revert_resolveCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_revert_resolveCmd.Flags().String("targets", "", "Path to a targets file")
	revision_revertCmd.AddCommand(revision_revert_resolveCmd)
}
