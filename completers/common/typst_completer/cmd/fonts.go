package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"os"
)

var fontsCmd = &cobra.Command{
	Use:   "fonts",
	Short: "Lists all discovered fonts in system and custom font paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fontsCmd).Standalone()

	fontsCmd.Flags().String("font-path", "", "Adds additional directories that are recursively searched for fonts.")
	fontsCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")
	fontsCmd.Flags().Bool("ignore-system-fonts", false, "Ensures system fonts won't be searched, unless explicitly included via `--font-path`")
	fontsCmd.Flags().Bool("variants", false, "Also lists style variants of each font family")

	rootCmd.AddCommand(fontsCmd)

	carapace.Gen(fontsCmd).FlagCompletion(carapace.ActionMap{
		"font-path": carapace.ActionDirectories().List(string(os.PathListSeparator)),
	})
}
