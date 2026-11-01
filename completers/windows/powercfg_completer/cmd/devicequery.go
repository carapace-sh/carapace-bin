package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicequeryCmd = &cobra.Command{
	Use:   "devicequery",
	Short: "query devices that meet certain criteria",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicequeryCmd).Standalone()
	rootCmd.AddCommand(devicequeryCmd)
}
