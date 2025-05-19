package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:     "unlink",
	Short:   "Disconnect the local project from another one",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkCmd).Standalone()

	unlinkCmd.Flags().BoolP("all", "A", false, "Unlink all workspaces belonging to the target project from the current one")
	rootCmd.AddCommand(unlinkCmd)

	// TODO positional completion
}
