package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var farm_listCmd = &cobra.Command{
	Use:     "list [options]",
	Short:   "List all existing farms",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(farm_listCmd).Standalone()

	farm_listCmd.Flags().String("format", "", "Format farm output using Go template")
	farmCmd.AddCommand(farm_listCmd)
}
