package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfHostedRunner_doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnose a deployed runner",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfHostedRunner_doctorCmd).Standalone()

	selfHostedRunnerCmd.AddCommand(selfHostedRunner_doctorCmd)
}
