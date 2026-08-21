package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var credentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(credentialCmd).Standalone()

	rootCmd.AddCommand(credentialCmd)
}
