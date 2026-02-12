package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "*Experimental* Start the background daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemonCmd).Standalone()

	daemonCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(daemonCmd)
}
