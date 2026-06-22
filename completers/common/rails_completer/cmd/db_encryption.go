package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbEncryptionCmd = &cobra.Command{
	Use:   "encryption",
	Short: "Database encryption tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbEncryptionCmd).Standalone()
	dbEncryptionCmd.Flags().BoolP("help", "h", false, "Show help")
}
