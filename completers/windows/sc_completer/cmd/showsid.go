package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showsidCmd = &cobra.Command{
	Use:   "showsid",
	Short: "display the service SID string for a name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showsidCmd).Standalone()
	rootCmd.AddCommand(showsidCmd)
}
