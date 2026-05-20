package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/typst_completer/cmd/common"
	"github.com/spf13/cobra"
)

var fontsCmd = &cobra.Command{
	Use:   "fonts",
	Short: "Lists all discovered fonts in system and custom font paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fontsCmd).Standalone()

	fontsCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")
	fontsCmd.Flags().Bool("variants", false, "Also lists style variants of each font family")
	common.AddFontFlags(fontsCmd)
	rootCmd.AddCommand(fontsCmd)
}
