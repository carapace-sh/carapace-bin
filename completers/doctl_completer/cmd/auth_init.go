package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var auth_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize doctl to use a specific account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_initCmd).Standalone()
	authCmd.AddCommand(auth_initCmd)
}
