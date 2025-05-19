package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_pushCmd = &cobra.Command{
	Use:   "push [OPTIONS] MANIFEST_LIST",
	Short: "Push a manifest list to a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_pushCmd).Standalone()

	manifest_pushCmd.Flags().Bool("insecure", false, "Allow push to an insecure registry")
	manifest_pushCmd.Flags().BoolP("purge", "p", false, "Remove the local manifest list after push")
	manifestCmd.AddCommand(manifest_pushCmd)

	// TODO completion
}
