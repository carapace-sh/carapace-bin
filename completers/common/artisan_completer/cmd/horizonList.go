package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonListCmd = &cobra.Command{
	Use:   "horizon:list",
	Short: "List all of the deployed machines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonListCmd).Standalone()

	rootCmd.AddCommand(horizonListCmd)
}
