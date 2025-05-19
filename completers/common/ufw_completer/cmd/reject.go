package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rejectCmd = &cobra.Command{
	Use:   "reject",
	Short: "add reject rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rejectCmd).Standalone()

	rootCmd.AddCommand(rejectCmd)
}
