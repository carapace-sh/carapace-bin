package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "manage tokenized interface identifiers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tokenCmd).Standalone()

	rootCmd.AddCommand(tokenCmd)
}
