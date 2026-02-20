package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chooseFontsCmd = &cobra.Command{
	Use:   "choose-fonts",
	Short: "Choose the fonts used in kitty",
}

func init() {
	rootCmd.AddCommand(chooseFontsCmd)
	carapace.Gen(chooseFontsCmd).Standalone()

	chooseFontsCmd.Flags().String("config-file-name", "", "The name or path to the config file to edit. Relative paths are interpreted with respect to the kitty config directory. By default the kitty config file, kitty.conf is edited. This is most useful if you add include themes.conf to your kitty.conf and then have the kitten operate only on themes.conf, allowing kitty.conf to remain unchanged.")
	chooseFontsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	chooseFontsCmd.Flags().String("reload-in", "", "By default, this kitten will signal only the parent kitty instance it is running in to reload its config, after making changes. Use this option to instead either not reload the config at all or in all running kitty instances.")

	carapace.Gen(chooseFontsCmd).FlagCompletion(carapace.ActionMap{
		"config-file-name": carapace.ActionFiles("~/.config/kitty"),
		"reload-in":        carapace.ActionValues("parent", "all", "none"),
	})

	carapace.Gen(chooseFontsCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
