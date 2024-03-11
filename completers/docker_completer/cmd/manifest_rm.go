package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_rmCmd = &cobra.Command{
	Use:   "rm MANIFEST_LIST [MANIFEST_LIST...]",
	Short: "Delete one or more manifest lists from local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_rmCmd).Standalone()

	manifestCmd.AddCommand(manifest_rmCmd)

	// TODO completion
}
