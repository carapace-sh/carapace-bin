package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:     "format",
	Short:   "Format TOML documents",
	Aliases: []string{"fmt"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().String("cache-path", "", "Set a cache path")
	formatCmd.Flags().Bool("check", false, "Dry-run and report any files that are not correctly formatted")
	formatCmd.Flags().StringP("config", "c", "", "Path to the Taplo configuration file")
	formatCmd.Flags().Bool("diff", false, "Print the differences in patch formatting to `stdout`")
	formatCmd.Flags().BoolP("force", "f", false, "Ignore syntax errors and force formatting")
	formatCmd.Flags().Bool("no-auto-config", false, "Do not search for a configuration file")
	formatCmd.Flags().StringSliceP("option", "o", nil, "A formatter option given as a \"key=value\", can be set multiple times")
	formatCmd.Flags().String("stdin-filepath", "", "A path to the file that the Taplo CLI will treat like stdin")
	rootCmd.AddCommand(formatCmd)

	carapace.Gen(formatCmd).FlagCompletion(carapace.ActionMap{
		"cache-path":     carapace.ActionDirectories(),
		"config":         carapace.ActionFiles(),
		"stdin-filepath": carapace.ActionFiles(),
	})

	carapace.Gen(formatCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
