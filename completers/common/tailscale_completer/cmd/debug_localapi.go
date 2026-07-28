package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_localapiCmd = &cobra.Command{
	Use:   "localapi",
	Short: "Call a LocalAPI method directly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_localapiCmd).Standalone()

	debug_localapiCmd.Flags().Bool("v", false, "verbose HTTP dump")
	debugCmd.AddCommand(debug_localapiCmd)
}
