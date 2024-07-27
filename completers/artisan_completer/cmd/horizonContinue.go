package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonContinueCmd = &cobra.Command{
	Use:   "horizon:continue",
	Short: "Instruct the master supervisor to continue processing jobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonContinueCmd).Standalone()

	rootCmd.AddCommand(horizonContinueCmd)
}
