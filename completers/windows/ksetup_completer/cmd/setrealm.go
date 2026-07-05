package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setrealmCmd = &cobra.Command{
	Use:   "setrealm",
	Short: "set the name of a Kerberos realm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setrealmCmd).Standalone()
	rootCmd.AddCommand(setrealmCmd)
}
