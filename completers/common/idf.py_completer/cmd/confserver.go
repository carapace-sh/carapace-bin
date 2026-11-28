package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var confserverCmd = &cobra.Command{
	Use:   "confserver",
	Short: "Run JSON configuration server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(confserverCmd).Standalone()
	confserverCmd.Flags().String("buffer-size", "", "Set the buffer size (in KB)")
	rootCmd.AddCommand(confserverCmd)
}
