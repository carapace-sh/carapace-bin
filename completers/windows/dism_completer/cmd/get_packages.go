package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetPackagesCmd = &cobra.Command{
	Use:   "Get-Packages",
	Short: "display all packages in an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetPackagesCmd).Standalone()
	rootCmd.AddCommand(GetPackagesCmd)
}
