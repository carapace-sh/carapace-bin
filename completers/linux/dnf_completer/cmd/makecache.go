package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makecacheCmd = &cobra.Command{
	Use:     "makecache",
	Aliases: []string{"mc"},
	Short:   "generate the metadata cache",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makecacheCmd).Standalone()

	rootCmd.AddCommand(makecacheCmd)
}
