package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codamaCmd = &cobra.Command{
	Use:   "codama",
	Short: "Codama IDL integration commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codamaCmd).Standalone()

	codamaCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(codamaCmd)
}
