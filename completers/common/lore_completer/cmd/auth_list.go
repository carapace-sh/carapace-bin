package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored authentication identities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_listCmd).Standalone()

	auth_listCmd.Flags().BoolP("help", "h", false, "Print help")
	auth_listCmd.Flags().Bool("with-token", false, "Include cached tokens in the output")
	authCmd.AddCommand(auth_listCmd)
}
