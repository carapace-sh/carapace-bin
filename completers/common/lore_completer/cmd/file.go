package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "File commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fileCmd).Standalone()

	fileCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(fileCmd)
}
