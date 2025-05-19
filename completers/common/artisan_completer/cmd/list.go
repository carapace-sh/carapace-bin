package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [--raw] [--format FORMAT] [--short] [--] [<namespace>]",
	Short: "List commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("format", "", "The output format (txt, xml, json, or md)")
	listCmd.Flags().Bool("raw", false, "To output raw command list")
	listCmd.Flags().Bool("short", false, "To skip describing commands' arguments")
	rootCmd.AddCommand(listCmd)
}
