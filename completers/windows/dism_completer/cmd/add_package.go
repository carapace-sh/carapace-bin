package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var AddPackageCmd = &cobra.Command{
	Use:   "Add-Package",
	Short: "add a package to an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(AddPackageCmd).Standalone()
	rootCmd.AddCommand(AddPackageCmd)
}
