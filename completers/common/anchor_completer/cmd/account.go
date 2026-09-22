package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Fetch and deserialize an account using the IDL provided",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accountCmd).Standalone()

	accountCmd.Flags().BoolP("help", "h", false, "Print help")
	accountCmd.Flags().String("idl", "", "Path of IDL to use (defaults to workspace IDL)")
	rootCmd.AddCommand(accountCmd)
}
