package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var betaCmd = &cobra.Command{
	Use:    "beta",
	Short:  "Experimental features",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(betaCmd).Standalone()

	rootCmd.AddCommand(betaCmd)
}
