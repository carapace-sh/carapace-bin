package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qmanagedaccountCmd = &cobra.Command{
	Use:   "qmanagedaccount",
	Short: "query whether a service uses an LSA-managed account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qmanagedaccountCmd).Standalone()
	rootCmd.AddCommand(qmanagedaccountCmd)
}
