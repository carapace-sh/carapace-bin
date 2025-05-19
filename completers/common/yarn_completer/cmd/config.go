package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Display the current configuration",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	configCmd.Flags().BoolP("verbose", "v", false, "Print the setting description on top of the regular key/value information")
	configCmd.Flags().Bool("why", false, "Print the reason why a setting is set a particular way")
	rootCmd.AddCommand(configCmd)
}
