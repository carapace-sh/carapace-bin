package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bundle_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List refs recorded in a bundle header",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_listCmd).Standalone()

	bundle_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bundle_listCmd.Flags().StringP("limit", "n", "", "Maximum number of rows to list")
	bundle_listCmd.Flags().String("max-count", "", "Maximum number of rows to list")
	bundleCmd.AddCommand(bundle_listCmd)

	carapace.Gen(bundle_listCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
