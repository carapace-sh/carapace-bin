package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstalCmd = &cobra.Command{
	Use:   "uninstal",
	Short: "Uninstall a <formula> or <cask>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstalCmd).Standalone()

	uninstalCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	uninstalCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	uninstalCmd.Flags().BoolP("force", "f", false, "Delete all installed versions of <formula>. Uninstall even if <cask> is not installed, overwrite existing files and ignore errors when removing files.")
	uninstalCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	uninstalCmd.Flags().BoolP("help", "h", false, "Show this message.")
	uninstalCmd.Flags().Bool("ignore-dependencies", false, "Don't fail uninstall, even if <formula> is a dependency of any installed formulae.")
	uninstalCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	uninstalCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	uninstalCmd.Flags().Bool("zap", false, "Remove all files associated with a <cask>. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(uninstalCmd)

	carapace.Gen(uninstalCmd).PositionalAnyCompletion(
		action.ActionList(uninstalCmd).FilterArgs(),
	)
}
