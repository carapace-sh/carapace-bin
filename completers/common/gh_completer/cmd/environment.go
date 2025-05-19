package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentCmd = &cobra.Command{
	Use:    "environment",
	Short:  "Environment variables that can be used with gh",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentCmd).Standalone()

	rootCmd.AddCommand(environmentCmd)
}
