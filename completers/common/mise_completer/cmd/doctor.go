package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:     "doctor",
	Short:   "Check mise installation and environment",
	Aliases: []string{"dr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().Bool("full", false, "Run full diagnostics")
	rootCmd.AddCommand(doctorCmd)
}
