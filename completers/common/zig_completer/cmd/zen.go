package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var zenCmd = &cobra.Command{
	Use:   "zen",
	Short: "Print Zen of Zig and exit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(zenCmd).Standalone()

	rootCmd.AddCommand(zenCmd)

	zenCmd.Flags().BoolP("help", "h", false, "Print help")
}
