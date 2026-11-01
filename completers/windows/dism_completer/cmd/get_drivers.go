package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetDriversCmd = &cobra.Command{
	Use:   "Get-Drivers",
	Short: "list driver packages in an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetDriversCmd).Standalone()
	rootCmd.AddCommand(GetDriversCmd)
}
