package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var suspendThenHibernateCmd = &cobra.Command{
	Use:     "suspend-then-hibernate",
	Short:   "Suspend the system, wake after a period of",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspendThenHibernateCmd).Standalone()

	rootCmd.AddCommand(suspendThenHibernateCmd)
}
