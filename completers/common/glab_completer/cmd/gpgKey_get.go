package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gpgKey_getCmd = &cobra.Command{
	Use:   "get <key-id>",
	Short: "Returns a single GPG key specified by the ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKey_getCmd).Standalone()

	gpgKey_getCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	gpgKey_getCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	gpgKeyCmd.AddCommand(gpgKey_getCmd)
}
