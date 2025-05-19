package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var listFontsCmd = &cobra.Command{
	Use:   "+list-fonts",
	Short: "list all the available fonts for Ghostty",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listFontsCmd).Standalone()

	listFontsCmd.Flags().Bool("bold", false, "filter results to bold style")
	listFontsCmd.Flags().String("family", "", "filter results to a specific family")
	listFontsCmd.Flags().Bool("help", false, "show help")
	listFontsCmd.Flags().Bool("italic", false, "filter results to italic style")
	rootCmd.AddCommand(listFontsCmd)

	carapace.Gen(listFontsCmd).FlagCompletion(carapace.ActionMap{
		"family": os.ActionFontFamilies(),
	})
}
