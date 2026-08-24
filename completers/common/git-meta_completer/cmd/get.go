package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get string metadata value(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolP("help", "h", false, "Print help")
	getCmd.Flags().Bool("json", false, "Output as JSON")
	getCmd.Flags().Bool("with-authorship", false, "Include authorship info (requires --json)")
	rootCmd.AddCommand(getCmd)
}
