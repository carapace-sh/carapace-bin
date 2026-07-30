package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetPackageInfoCmd = &cobra.Command{
	Use:   "Get-PackageInfo",
	Short: "display information about a specific package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetPackageInfoCmd).Standalone()
	rootCmd.AddCommand(GetPackageInfoCmd)
}
