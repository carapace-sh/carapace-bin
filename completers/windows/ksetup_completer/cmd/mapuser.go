package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mapuserCmd = &cobra.Command{
	Use:   "mapuser",
	Short: "map a user account to a Kerberos principal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mapuserCmd).Standalone()
	rootCmd.AddCommand(mapuserCmd)
}
