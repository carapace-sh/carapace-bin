package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_rmCmd = &cobra.Command{
	Use:   "rm [options] LIST [LIST...]",
	Short: "Remove manifest list or image index from local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_rmCmd).Standalone()

	manifest_rmCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified manifest is missing")
	manifestCmd.AddCommand(manifest_rmCmd)
}
