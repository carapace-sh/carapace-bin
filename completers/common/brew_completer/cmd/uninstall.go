package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Uninstall a <formula> or <cask>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	uninstallCmd.Flags().Bool("debug", false, "Display any debugging information.")
	uninstallCmd.Flags().Bool("force", false, "Delete all installed versions of <formula>. Uninstall even if <cask> is not installed, overwrite existing files and ignore errors when removing files.")
	uninstallCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	uninstallCmd.Flags().Bool("help", false, "Show this message.")
	uninstallCmd.Flags().Bool("ignore-dependencies", false, "Don't fail uninstall, even if <formula> is a dependency of any installed formulae.")
	uninstallCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	uninstallCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	uninstallCmd.Flags().Bool("zap", false, "Remove all files associated with a <cask>. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		action.ActionList(uninstallCmd).FilterArgs(),
	)
}
