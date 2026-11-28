package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var refreshConfigCmd = &cobra.Command{
	Use:   "refresh-config",
	Short: "Resolve conflicts in default values of config options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refreshConfigCmd).Standalone()
	refreshConfigCmd.Flags().String("policy", "", "Policy for handling defaults in the configuration")
	rootCmd.AddCommand(refreshConfigCmd)
	carapace.Gen(refreshConfigCmd).FlagCompletion(carapace.ActionMap{
		"policy": carapace.ActionValues("interactive", "kconfig", "sdkconfig"),
	})
}
