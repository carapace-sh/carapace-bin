package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "Remove all routes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flushCmd).Standalone()
	flushCmd.Flags().Bool("inet", false, "Flush IPv4 routes only")
	flushCmd.Flags().Bool("inet6", false, "Flush IPv6 routes only")
	rootCmd.AddCommand(flushCmd)
}
