package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_policiesCmd = &cobra.Command{
	Use:   "policies",
	Short: "Will show the default branch policies by default",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_policiesCmd).Standalone()

	open_policiesCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_policiesCmd)
}
