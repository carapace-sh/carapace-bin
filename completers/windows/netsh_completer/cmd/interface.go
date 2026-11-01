package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "interface configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(interfaceCmd).Standalone()
	rootCmd.AddCommand(interfaceCmd)
}
