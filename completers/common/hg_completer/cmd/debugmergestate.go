package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugmergestateCmd = &cobra.Command{
	Use:    "debugmergestate",
	Short:  "print merge state",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugmergestateCmd).Standalone()

	debugmergestateCmd.Flags().String("style", "", "display using template map file (DEPRECATED)")
	debugmergestateCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugmergestateCmd)
}
