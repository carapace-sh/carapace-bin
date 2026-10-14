package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pivCmd = &cobra.Command{
	Use:    "piv",
	Short:  "PIV commands.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pivCmd).Standalone()

	rootCmd.AddCommand(pivCmd)
}
