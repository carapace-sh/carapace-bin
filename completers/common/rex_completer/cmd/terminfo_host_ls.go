package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminfoHostLsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List the hosts recorded as having Rex's terminfo entry",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoHostLsCmd).Standalone()

	terminfoHostLsCmd.Flags().Bool("json", false, "output the list as JSON")

	terminfoHostCmd.AddCommand(terminfoHostLsCmd)
}
