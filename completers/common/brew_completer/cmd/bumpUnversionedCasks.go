package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var bumpUnversionedCasksCmd = &cobra.Command{
	Use:     "bump-unversioned-casks",
	Short:   "Check all casks with unversioned URLs in a given <tap> for updates",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpUnversionedCasksCmd).Standalone()

	bumpUnversionedCasksCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	bumpUnversionedCasksCmd.Flags().BoolP("dry-run", "n", false, "Do everything except caching state and opening pull requests.")
	bumpUnversionedCasksCmd.Flags().BoolP("help", "h", false, "Show this message.")
	bumpUnversionedCasksCmd.Flags().String("limit", "", "Maximum runtime in minutes.")
	bumpUnversionedCasksCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	bumpUnversionedCasksCmd.Flags().String("state-file", "", "File for caching state.")
	bumpUnversionedCasksCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(bumpUnversionedCasksCmd)

	carapace.Gen(bumpUnversionedCasksCmd).FlagCompletion(carapace.ActionMap{
		"state-file": carapace.ActionFiles(),
	})

	carapace.Gen(bumpUnversionedCasksCmd).PositionalAnyCompletion(
		carapace.Batch(
			brew.ActionAllCasks(),
			brew.ActionAllFormulae(),
		).ToA().FilterArgs(),
	)
}
