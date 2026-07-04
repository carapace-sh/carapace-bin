package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "Delete system logging data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eraseCmd).Standalone()

	eraseCmd.Flags().Bool("all", false, "Erase all log data")
	eraseCmd.Flags().Bool("ttl", false, "Erase log data marked with a time-to-live")
	rootCmd.AddCommand(eraseCmd)
}
