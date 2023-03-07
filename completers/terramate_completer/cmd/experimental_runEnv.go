package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var experimental_runEnvCmd = &cobra.Command{
	Use:   "run-env",
	Short: "List run environment variables for all stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_runEnvCmd).Standalone()

	experimentalCmd.AddCommand(experimental_runEnvCmd)
}
