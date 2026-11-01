package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wakesourcesCmd = &cobra.Command{
	Use:   "wakesources",
	Short: "list devices that wake the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wakesourcesCmd).Standalone()
	rootCmd.AddCommand(wakesourcesCmd)
}
