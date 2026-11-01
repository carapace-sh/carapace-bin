package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedserviceaccountCmd = &cobra.Command{
	Use:   "managedserviceaccount",
	Short: "mark a service account as managed by LSA",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedserviceaccountCmd).Standalone()
	rootCmd.AddCommand(managedserviceaccountCmd)
}
