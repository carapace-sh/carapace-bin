package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Diagnoses the system to identify common errors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnoseCmd).Standalone()

	rootCmd.AddCommand(diagnoseCmd)
}
