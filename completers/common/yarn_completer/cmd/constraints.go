package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var constraintsCmd = &cobra.Command{
	Use:     "constraints",
	Short:   "manage constraints",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(constraintsCmd).Standalone()

	rootCmd.AddCommand(constraintsCmd)
}
