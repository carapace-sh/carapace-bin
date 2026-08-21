package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var access_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_tokenCmd).Standalone()

	accessCmd.AddCommand(access_tokenCmd)
}
