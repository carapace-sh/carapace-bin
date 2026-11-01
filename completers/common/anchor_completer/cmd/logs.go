package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Stream transaction logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()

	logsCmd.Flags().StringSlice("address", nil, "Addresses to filter logs by")
	logsCmd.Flags().BoolP("help", "h", false, "Print help")
	logsCmd.Flags().Bool("include-votes", false, "Include vote transactions when monitoring all transactions")
	rootCmd.AddCommand(logsCmd)
}
