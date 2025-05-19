package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hibernateCmd = &cobra.Command{
	Use:     "hibernate",
	Short:   "Hibernate the system",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hibernateCmd).Standalone()

	rootCmd.AddCommand(hibernateCmd)
}
