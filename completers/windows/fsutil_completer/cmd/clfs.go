package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clfsCmd = &cobra.Command{
	Use:   "clfs",
	Short: "manage Common Log File System logfiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clfsCmd).Standalone()
	rootCmd.AddCommand(clfsCmd)
}
