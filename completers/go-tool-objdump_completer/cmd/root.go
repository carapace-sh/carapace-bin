package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-objdump",
	Short: "Objdump disassembles executable files",
	Long:  "https://pkg.go.dev/cmd/objdump",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "print Go code alongside assembly")
	rootCmd.Flags().BoolS("gnu", "gnu", false, "print GNU assembly next to Go assembly")
	rootCmd.Flags().StringS("s", "s", "", "only dump symbols matching this regexp")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
