package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "manage zones",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(zoneCmd).Standalone()
	rootCmd.AddCommand(zoneCmd)
}
