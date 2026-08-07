package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deployKey_getCmd = &cobra.Command{
	Use:   "get <key-id>",
	Short: "Returns a single deploy key specified by the ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployKey_getCmd).Standalone()

	deployKey_getCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	deployKey_getCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	deployKeyCmd.AddCommand(deployKey_getCmd)

	// TODO positional completion
}
