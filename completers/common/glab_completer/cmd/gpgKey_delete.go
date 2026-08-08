package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gpgKey_deleteCmd = &cobra.Command{
	Use:   "delete <key-id>",
	Short: "Deletes a single GPG key specified by the ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKey_deleteCmd).Standalone()

	gpgKeyCmd.AddCommand(gpgKey_deleteCmd)
}
