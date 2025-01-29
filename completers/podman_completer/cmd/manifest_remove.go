package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_removeCmd = &cobra.Command{
	Use:   "remove LIST DIGEST",
	Short: "Remove an item from a manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_removeCmd).Standalone()

	manifestCmd.AddCommand(manifest_removeCmd)
}
