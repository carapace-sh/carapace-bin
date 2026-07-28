package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_dialTypesCmd = &cobra.Command{
	Use:   "dial-types",
	Short: "Print debug information about connecting to a given host or IP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_dialTypesCmd).Standalone()

	debug_dialTypesCmd.Flags().String("network", "", "network type (e.g. tcp, udp)")
	debugCmd.AddCommand(debug_dialTypesCmd)

	carapace.Gen(debug_dialTypesCmd).FlagCompletion(carapace.ActionMap{
		"network": carapace.ActionValues("tcp", "udp"),
	})
}
