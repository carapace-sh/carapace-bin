package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var narCmd = &cobra.Command{
	Use:     "nar",
	Short:   "create or inspect NAR files",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(narCmd).Standalone()

	rootCmd.AddCommand(narCmd)
}
