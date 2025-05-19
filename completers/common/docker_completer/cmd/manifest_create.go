package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_createCmd = &cobra.Command{
	Use:   "create MANIFEST_LIST MANIFEST [MANIFEST...]",
	Short: "Create a local manifest list for annotating and pushing to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_createCmd).Standalone()

	manifest_createCmd.Flags().BoolP("amend", "a", false, "Amend an existing manifest list")
	manifest_createCmd.Flags().Bool("insecure", false, "Allow communication with an insecure registry")
	manifestCmd.AddCommand(manifest_createCmd)

	// TODO completion
}
