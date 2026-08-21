package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var availabilityCmd = &cobra.Command{
	Use:   "availability",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(availabilityCmd).Standalone()

	rootCmd.AddCommand(availabilityCmd)
}
