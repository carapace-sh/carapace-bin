package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "List or describe the available configuration options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.Flags().String("output-format", "text", "Output format")
	rootCmd.AddCommand(configCmd)
}
