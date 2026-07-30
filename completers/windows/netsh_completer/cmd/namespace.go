package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "namespace configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespaceCmd).Standalone()
	rootCmd.AddCommand(namespaceCmd)
}
