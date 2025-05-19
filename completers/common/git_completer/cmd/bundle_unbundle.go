package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bundle_unbundleCmd = &cobra.Command{
	Use:   "unbundle",
	Short: "passes the objects in the bundle for storage in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_unbundleCmd).Standalone()

	bundleCmd.AddCommand(bundle_unbundleCmd)

	carapace.Gen(bundle_unbundleCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(bundle_unbundleCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionBundleHeads(c.Args[0])
		}),
	)
}
