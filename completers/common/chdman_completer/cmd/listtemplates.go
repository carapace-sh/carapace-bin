package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listtemplatesCmd = &cobra.Command{
	Use:   "listtemplates",
	Short: "list hard disk templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listtemplatesCmd).Standalone()

	rootCmd.AddCommand(listtemplatesCmd)
}
