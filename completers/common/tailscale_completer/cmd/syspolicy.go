package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syspolicyCmd = &cobra.Command{
	Use:   "syspolicy",
	Short: "Diagnose the MDM and system policy configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syspolicyCmd).Standalone()

	rootCmd.AddCommand(syspolicyCmd)
}
