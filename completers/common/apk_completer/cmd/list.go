package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List packages matching a pattern or other criteria",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "querying package information",
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("available", "a", false, "Consider only available packages")
	listCmd.Flags().BoolP("depends", "d", false, "List packages by dependency")
	listCmd.Flags().BoolP("installed", "I", false, "Consider only installed packages")
	listCmd.Flags().BoolP("origin", "o", false, "List packages by origin")
	listCmd.Flags().BoolP("orphaned", "O", false, "Consider only orphaned packages")
	listCmd.Flags().BoolP("providers", "P", false, "List packages by provider")
	listCmd.Flags().BoolP("upgradable", "u", false, "Consider only upgradable packages")
	listCmd.Flags().Bool("upgradeable", false, "Consider only upgradable packages")
	common.AddGlobalFlags(listCmd)
	common.AddSourceFlags(listCmd)
	rootCmd.AddCommand(listCmd)
}
