package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var haspersonalrecoverykeyCmd = &cobra.Command{
	Use:   "haspersonalrecoverykey",
	Short: "Check if a personal recovery key has been set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(haspersonalrecoverykeyCmd).Standalone()
	rootCmd.AddCommand(haspersonalrecoverykeyCmd)
}
