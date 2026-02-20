package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themesCmd = &cobra.Command{
	Use:   "themes",
	Short: "Manage kitty color schemes easily",
}

func init() {
	rootCmd.AddCommand(themesCmd)
	carapace.Gen(themesCmd).Standalone()

	themesCmd.Flags().Float64("cache-age", 0, "Check for new themes only after the specified number of days. A value of zero will always check for new themes. A negative value will never check for new themes, instead raising an error if a local copy of the themes is not available.")
	themesCmd.Flags().String("config-file-name", "", "The name or path to the config file to edit. Relative paths are interpreted with respect to the kitty config directory. By default the kitty config file, kitty.conf is edited. This is most useful if you add include themes.conf to your kitty.conf and then have the kitten operate only on themes.conf, allowing kitty.conf to remain unchanged.")
	themesCmd.Flags().Bool("dump-theme", false, "When running non-interactively, dump the specified theme to STDOUT instead of changing kitty.conf.")
	themesCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	themesCmd.Flags().String("reload-in", "", "By default, this kitten will signal only the parent kitty instance it is running in to reload its config, after making changes. Use this option to instead either not reload the config at all or in all running kitty instances.")

	carapace.Gen(themesCmd).FlagCompletion(carapace.ActionMap{
		"reload-in":        carapace.ActionValues("parent", "all", "none"),
		"config-file-name": carapace.ActionFiles("~/.config/kitty"),
	})
}
