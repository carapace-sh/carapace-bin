package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var driversCmd = &cobra.Command{
	Use:   "drivers [flags]",
	Short: "List drivers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(driversCmd)

	driversCmd.Flags().StringP("format", "f", "plain", "Change output format")
	driversCmd.Flags().Bool("help", false, "Show help message and exit")

	carapace.Gen(driversCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "simple"),
	})

	rootCmd.AddCommand(driversCmd)
}
