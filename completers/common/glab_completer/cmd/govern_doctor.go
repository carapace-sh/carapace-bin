package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var govern_doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose AI agent governance configuration, authentication, and hook setup. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(govern_doctorCmd).Standalone()

	governCmd.AddCommand(govern_doctorCmd)
}
