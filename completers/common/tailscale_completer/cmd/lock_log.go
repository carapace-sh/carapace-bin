package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_logCmd = &cobra.Command{
	Use:   "log",
	Short: "List changes applied to tailnet lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_logCmd).Standalone()

	lock_logCmd.Flags().Bool("json", false, "output in JSON format")
	lock_logCmd.Flags().String("limit", "", "max number of updates to list")
	lockCmd.AddCommand(lock_logCmd)
}
