package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdsetCmd = &cobra.Command{
	Use:   "sdset",
	Short: "set a service's security descriptor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdsetCmd).Standalone()
	rootCmd.AddCommand(sdsetCmd)
}
