package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var delrealmCmd = &cobra.Command{
	Use:   "delrealm",
	Short: "delete the Kerberos realm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delrealmCmd).Standalone()
	rootCmd.AddCommand(delrealmCmd)
}
