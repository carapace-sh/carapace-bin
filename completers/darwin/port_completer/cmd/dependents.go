package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dependentsCmd = &cobra.Command{
	Use:   "dependents",
	Short: "List installed ports that depend on a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependentsCmd).Standalone()
	rootCmd.AddCommand(dependentsCmd)
}
