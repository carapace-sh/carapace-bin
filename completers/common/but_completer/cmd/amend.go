package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Amend one or more file changes into a specific commit and rebases any dependent commits",
	Run:   func(cmd *cobra.Command, args []string) {},
	GroupID: "editing commits",
}

func init() {
	carapace.Gen(amendCmd).Standalone()

	amendCmd.Flags().StringSliceP("changes", "p", nil, "Uncommitted file or hunk CLI IDs to amend into the commit")
	amendCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(amendCmd)
}
