package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var localCmd = &cobra.Command{
	Use:   "local",
	Short: "enable or disable the Maintain Objects List global flag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localCmd).Standalone()
	rootCmd.AddCommand(localCmd)
}
