package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var usingrecoverykeyCmd = &cobra.Command{
	Use:   "usingrecoverykey",
	Short: "Check if a recovery key is currently in use",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usingrecoverykeyCmd).Standalone()
	rootCmd.AddCommand(usingrecoverykeyCmd)
}
