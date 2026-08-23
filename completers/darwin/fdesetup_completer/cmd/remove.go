package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove user from FileVault",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("quiet", false, "No status during operation")
	removeCmd.Flags().String("user", "", "Short user name")
	removeCmd.Flags().String("uuid", "", "User UUID")
	removeCmd.Flags().Bool("verbose", false, "Enable verbose mode")

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"user": os.ActionUsers(),
	})
}
