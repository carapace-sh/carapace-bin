package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unmapuserCmd = &cobra.Command{
	Use:   "unmapuser",
	Short: "unmap a user account from a Kerberos principal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmapuserCmd).Standalone()
	rootCmd.AddCommand(unmapuserCmd)
}
