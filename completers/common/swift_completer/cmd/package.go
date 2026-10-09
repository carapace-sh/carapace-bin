package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Perform operations on Swift packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packageCmd).Standalone()
	packageCmd.Flags().SetInterspersed(false)

	common.AddPackageFlags(packageCmd)

	packageCmd.Flags().BoolP("help", "h", false, "Show help information")
	packageCmd.Flags().Bool("version", false, "Show the version")

	rootCmd.AddCommand(packageCmd)
}
