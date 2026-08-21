package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var implied_roleCmd = &cobra.Command{
	Use:   "role",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(implied_roleCmd).Standalone()

	impliedCmd.AddCommand(implied_roleCmd)
}
