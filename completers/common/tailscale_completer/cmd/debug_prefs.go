package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_prefsCmd = &cobra.Command{
	Use:   "prefs",
	Short: "Print prefs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_prefsCmd).Standalone()

	debug_prefsCmd.Flags().Bool("pretty", false, "pretty-print output")
	debugCmd.AddCommand(debug_prefsCmd)
}
