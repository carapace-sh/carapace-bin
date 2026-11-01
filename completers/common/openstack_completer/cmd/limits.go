package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var limitsCmd = &cobra.Command{
	Use:   "limits",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(limitsCmd).Standalone()

	rootCmd.AddCommand(limitsCmd)
}
