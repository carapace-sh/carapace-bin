package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var patchRemoveCmd = &cobra.Command{
	Use:   "patch-remove",
	Short: "Remove existing patch files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchRemoveCmd).Standalone()

	patchRemoveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(patchRemoveCmd)
}
