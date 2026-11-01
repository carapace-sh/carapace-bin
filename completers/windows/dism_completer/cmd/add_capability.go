package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var AddCapabilityCmd = &cobra.Command{
	Use:   "Add-Capability",
	Short: "add a capability to an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(AddCapabilityCmd).Standalone()
	rootCmd.AddCommand(AddCapabilityCmd)
}
