package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Uninstall a <formula> or <cask>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	rmCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	rmCmd.Flags().BoolP("force", "f", false, "Delete all installed versions of <formula>. Uninstall even if <cask> is not installed, overwrite existing files and ignore errors when removing files.")
	rmCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	rmCmd.Flags().BoolP("help", "h", false, "Show this message.")
	rmCmd.Flags().Bool("ignore-dependencies", false, "Don't fail uninstall, even if <formula> is a dependency of any installed formulae.")
	rmCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	rmCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rmCmd.Flags().Bool("zap", false, "Remove all files associated with a <cask>. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalAnyCompletion(
		action.ActionList(rmCmd).FilterArgs(),
	)
}
