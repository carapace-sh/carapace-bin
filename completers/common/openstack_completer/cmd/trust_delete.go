package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trust_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete trust(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_deleteCmd).Standalone()

	trustCmd.AddCommand(trust_deleteCmd)
}
