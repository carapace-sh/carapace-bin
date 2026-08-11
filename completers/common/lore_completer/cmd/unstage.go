package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unstageCmd = &cobra.Command{
	Use:   "unstage",
	Short: "Unstage changes to a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unstageCmd).Standalone()

	unstageCmd.Flags().BoolP("help", "h", false, "Print help")
	unstageCmd.Flags().String("targets", "", "Path to a targets file")
	rootCmd.AddCommand(unstageCmd)
}
