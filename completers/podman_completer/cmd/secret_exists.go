package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_existsCmd = &cobra.Command{
	Use:   "exists SECRET",
	Short: "Check if a secret exists in local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_existsCmd).Standalone()

	secretCmd.AddCommand(secret_existsCmd)
}
