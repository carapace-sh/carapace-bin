package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkupCmd = &cobra.Command{
	Use:   "checkup",
	Short: "check for potential problems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkupCmd).Standalone()
	rootCmd.AddCommand(checkupCmd)
}
