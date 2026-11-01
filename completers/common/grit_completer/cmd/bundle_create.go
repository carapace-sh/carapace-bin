package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bundle_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Write a bundle file from revision specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_createCmd).Standalone()

	bundle_createCmd.Flags().Bool("all", false, "Include every ref tip in the bundle header (`HEAD`, all branches and tags)")
	bundle_createCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bundleCmd.AddCommand(bundle_createCmd)

	carapace.Gen(bundle_createCmd).PositionalCompletion(
		carapace.ActionFiles(),
		git.ActionRefRanges(git.RefOption{}.Default()),
	)
}
