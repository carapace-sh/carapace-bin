package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var local_ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(local_ipCmd).Standalone()

	localCmd.AddCommand(local_ipCmd)
}
