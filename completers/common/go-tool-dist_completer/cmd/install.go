package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install individual directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolS("v", "v", false, "verbosity")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
