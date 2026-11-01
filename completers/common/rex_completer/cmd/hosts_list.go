package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hostsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your hosts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostsListCmd).Standalone()

	hostsListCmd.Flags().Bool("json", false, "output the hosts as JSON")

	hostsCmd.AddCommand(hostsListCmd)
}
