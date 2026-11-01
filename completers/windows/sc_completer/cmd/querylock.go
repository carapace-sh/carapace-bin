package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var querylockCmd = &cobra.Command{
	Use:   "querylock",
	Short: "query the lock status for the service database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(querylockCmd).Standalone()
	rootCmd.AddCommand(querylockCmd)
}
