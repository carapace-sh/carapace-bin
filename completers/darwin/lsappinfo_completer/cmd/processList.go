package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var processListCmd = &cobra.Command{
	Use:   "processList",
	Short: "Show the application list in ascending ASN order",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processListCmd).Standalone()
	rootCmd.AddCommand(processListCmd)
}
