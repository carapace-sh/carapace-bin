package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a basic job from the command line",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(submitCmd).Standalone()
	rootCmd.AddCommand(submitCmd)
	carapace.Gen(submitCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
