package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/linux/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:     "fix",
	Short:   "Fix, reinstall or upgrade packages without modifying WORLD",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "system maintenance",
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	fixCmd.Flags().BoolP("depends", "d", false, "Also fix dependencies of specified packages")
	fixCmd.Flags().Bool("directory-permissions", false, "Reset all directory permissions")
	fixCmd.Flags().BoolP("reinstall", "r", false, "Reinstall packages (default)")
	fixCmd.Flags().BoolP("upgrade", "u", false, "Upgrade name PACKAGE if an upgrade exists")
	fixCmd.Flags().BoolP("xattr", "x", false, "Fix packages with broken xattrs")
	common.AddGlobalFlags(fixCmd)
	common.AddCommitFlags(fixCmd)
	rootCmd.AddCommand(fixCmd)

	carapace.Gen(fixCmd).PositionalAnyCompletion(
		action.ActionPackages(fixCmd, true).FilterArgs(),
	)
}
