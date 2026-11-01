package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfHostedRunner_setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Guided operator walkthrough",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfHostedRunner_setupCmd).Standalone()

	selfHostedRunnerCmd.AddCommand(selfHostedRunner_setupCmd)
}
