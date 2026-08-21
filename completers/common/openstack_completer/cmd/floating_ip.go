package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ipCmd).Standalone()

	floatingCmd.AddCommand(floating_ipCmd)
}
