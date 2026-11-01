package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_macos_defaultsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Manage macOS preferences such as Dock, Finder, and keyboard settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_macos_defaultsCmd).Standalone()

	bootstrap_macos_defaultsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_macosCmd.AddCommand(bootstrap_macos_defaultsCmd)
}
