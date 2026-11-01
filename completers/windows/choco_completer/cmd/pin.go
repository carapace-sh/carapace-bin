package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "pin a package to prevent upgrades",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinCmd).Standalone()
	rootCmd.AddCommand(pinCmd)
}
