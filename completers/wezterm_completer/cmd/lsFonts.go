package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsFontsCmd = &cobra.Command{
	Use:   "ls-fonts [OPTIONS]",
	Short: "Display information about fonts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsFontsCmd).Standalone()

	lsFontsCmd.Flags().String("codepoints", "", "Explain which fonts are used to render the specified unicode code point sequence")
	lsFontsCmd.Flags().BoolP("help", "h", false, "Print help")
	lsFontsCmd.Flags().Bool("list-system", false, "Whether to list all fonts available to the system")
	lsFontsCmd.Flags().Bool("rasterize-ascii", false, "Show rasterized glyphs for the text in --text or --codepoints using ascii blocks")
	lsFontsCmd.Flags().String("text", "", "Explain which fonts are used to render the supplied text string")
	rootCmd.AddCommand(lsFontsCmd)
}
