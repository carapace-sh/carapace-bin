package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hashInfoCmd = &cobra.Command{
	Use:   "hostinfo",
	Short: "Print host information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hashInfoCmd).Standalone()
	rootCmd.AddCommand(hashInfoCmd)
}
