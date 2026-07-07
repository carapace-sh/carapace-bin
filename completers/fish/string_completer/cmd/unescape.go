package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unescapeCmd = &cobra.Command{
	Use:   "unescape",
	Short: "Unescape strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unescapeCmd).Standalone()

	unescapeCmd.Flags().String("style", "", "unescape style")

	rootCmd.AddCommand(unescapeCmd)

	carapace.Gen(unescapeCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("script", "var", "url"),
	})
}
