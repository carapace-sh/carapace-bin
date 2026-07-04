package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change to the directory containing a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cdCmd).Standalone()
	rootCmd.AddCommand(cdCmd)
}
