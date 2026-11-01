package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RemovePackageCmd = &cobra.Command{
	Use:   "Remove-Package",
	Short: "remove a package from an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RemovePackageCmd).Standalone()
	rootCmd.AddCommand(RemovePackageCmd)
}
