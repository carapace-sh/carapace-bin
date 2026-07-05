package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logicaldiskCmd = &cobra.Command{
	Use:   "logicaldisk",
	Short: "logical disk management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logicaldiskCmd).Standalone()
	rootCmd.AddCommand(logicaldiskCmd)
}
