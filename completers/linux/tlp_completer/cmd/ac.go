package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acCmd = &cobra.Command{
	Use:   "ac",
	Short: "Apply profile for AC power and enter manual mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acCmd).Standalone()

	rootCmd.AddCommand(acCmd)
}
