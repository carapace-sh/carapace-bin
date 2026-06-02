package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clean_profileCmd = &cobra.Command{
	Use:   "profile [profile]",
	Short: "Clean a specific profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clean_profileCmd).Standalone()

	addCleanFlags(clean_profileCmd)

	carapace.Gen(clean_profileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	cleanCmd.AddCommand(clean_profileCmd)
}
