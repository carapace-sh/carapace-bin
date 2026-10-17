package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "List and change your hosts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostsCmd).Standalone()

	hostsCmd.Flags().Bool("json", false, "output the hosts as JSON")

	rootCmd.AddCommand(hostsCmd)
}
