package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var escapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "Escape strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(escapeCmd).Standalone()

	escapeCmd.Flags().BoolP("no-quoted", "n", false, "don't use simplifying quoted format")
	escapeCmd.Flags().String("style", "", "escape style")

	rootCmd.AddCommand(escapeCmd)

	carapace.Gen(escapeCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("script", "var", "url", "regex"),
	})
}
