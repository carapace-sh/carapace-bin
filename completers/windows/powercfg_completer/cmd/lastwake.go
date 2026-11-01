package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lastwakeCmd = &cobra.Command{
	Use:   "lastwake",
	Short: "report information about what woke the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lastwakeCmd).Standalone()
	rootCmd.AddCommand(lastwakeCmd)
}
