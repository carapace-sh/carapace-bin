package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scpCmd = &cobra.Command{
	Use:   "scp",
	Short: "copy files to or from a machine via SCP",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scpCmd).Standalone()
	rootCmd.AddCommand(scpCmd)
}
