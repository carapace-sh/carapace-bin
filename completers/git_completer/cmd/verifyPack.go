package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyPackCmd = &cobra.Command{
	Use:     "verify-pack",
	Short:   "Validate packed Git archive files",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(verifyPackCmd).Standalone()

	verifyPackCmd.Flags().BoolP("stat-only", "s", false, "only show the histogram of delta chain length")
	verifyPackCmd.Flags().BoolP("verbose", "v", false, "show the list of objects contained in the pack")
	rootCmd.AddCommand(verifyPackCmd)

	carapace.Gen(verifyPackCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".idx").FilterArgs(),
	)

	carapace.Gen(verifyPackCmd).DashAnyCompletion(
		carapace.ActionPositional(verifyPackCmd),
	)
}
