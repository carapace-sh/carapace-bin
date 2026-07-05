package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var localgroupCmd = &cobra.Command{
	Use:   "localgroup",
	Short: "manage local groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localgroupCmd).Standalone()
	rootCmd.AddCommand(localgroupCmd)
}
