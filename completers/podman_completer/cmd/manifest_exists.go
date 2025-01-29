package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_existsCmd = &cobra.Command{
	Use:   "exists MANIFEST",
	Short: "Check if a manifest list exists in local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_existsCmd).Standalone()

	manifestCmd.AddCommand(manifest_existsCmd)
}
