package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var user_passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_passwordCmd).Standalone()

	userCmd.AddCommand(user_passwordCmd)
}
