package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "Print which machine is active",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(activeCmd).Standalone()

	activeCmd.Flags().StringP("timeout", "t", "", "Timeout in seconds, default to 10s")
	rootCmd.AddCommand(activeCmd)
}
