package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var npmCmd = &cobra.Command{
	Use:     "npm",
	Short:   "manage npm",
	GroupID: "npm-related",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(npmCmd).Standalone()

	rootCmd.AddCommand(npmCmd)
}
