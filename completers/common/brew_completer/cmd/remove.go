package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Uninstall a <formula> or <cask>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	removeCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	removeCmd.Flags().BoolP("force", "f", false, "Delete all installed versions of <formula>. Uninstall even if <cask> is not installed, overwrite existing files and ignore errors when removing files.")
	removeCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	removeCmd.Flags().BoolP("help", "h", false, "Show this message.")
	removeCmd.Flags().Bool("ignore-dependencies", false, "Don't fail uninstall, even if <formula> is a dependency of any installed formulae.")
	removeCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	removeCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	removeCmd.Flags().Bool("zap", false, "Remove all files associated with a <cask>. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		action.ActionList(removeCmd).FilterArgs(),
	)
}
