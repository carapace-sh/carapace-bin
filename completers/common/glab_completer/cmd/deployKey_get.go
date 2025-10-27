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

	deployKeyCmd.AddCommand(deployKey_getCmd)

	// TODO positional completion
}
