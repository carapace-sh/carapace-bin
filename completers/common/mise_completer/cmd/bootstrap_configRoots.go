package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_configRootsCmd = &cobra.Command{
	Use:    "config-roots",
	Short:  "Show non-composed bootstrap declarations in each selected configuration root (deprecated)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_configRootsCmd).Standalone()

	bootstrap_configRootsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_configRootsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrapCmd.AddCommand(bootstrap_configRootsCmd)
}
