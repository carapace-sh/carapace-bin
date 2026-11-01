package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RemoveCapabilityCmd = &cobra.Command{
	Use:   "Remove-Capability",
	Short: "remove a capability from an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RemoveCapabilityCmd).Standalone()
	rootCmd.AddCommand(RemoveCapabilityCmd)
}
