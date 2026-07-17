package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listAvailableCmd = &cobra.Command{
	Use:   "list-available [toolchain]",
	Short: "List toolchains available for install",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listAvailableCmd).Standalone()

	listAvailableCmd.Flags().String("format", "text", "Output format (text|json)")
	listAvailableCmd.Flags().BoolP("help", "h", false, "Show help information")

	rootCmd.AddCommand(listAvailableCmd)

	carapace.Gen(listAvailableCmd).FlagCompletion(carapace.ActionMap{
		"format": actionFormat(),
	})
}
