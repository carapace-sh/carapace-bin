package cmd

import (
	"github.com/carapace-sh/carapace"
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
	rmCmd.Flags().Bool("debug", false, "Display any debugging information.")
	rmCmd.Flags().Bool("force", false, "Delete all installed versions of <formula>. Uninstall even if <cask> is not installed, overwrite existing files and ignore errors when removing files.")
	rmCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	rmCmd.Flags().Bool("help", false, "Show this message.")
	rmCmd.Flags().Bool("ignore-dependencies", false, "Don't fail uninstall, even if <formula> is a dependency of any installed formulae.")
	rmCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	rmCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rmCmd.Flags().Bool("zap", false, "Remove all files associated with a <cask>. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(rmCmd)
}
