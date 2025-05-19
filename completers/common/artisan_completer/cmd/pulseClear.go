package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pulseClearCmd = &cobra.Command{
	Use:     "pulse:clear [--type [TYPE]] [--force]",
	Short:   "Delete all Pulse data from storage",
	Aliases: []string{"pulse:purge"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulseClearCmd).Standalone()

	pulseClearCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	pulseClearCmd.Flags().String("type", "", "Only clear the specified type(s)")
	rootCmd.AddCommand(pulseClearCmd)
}
