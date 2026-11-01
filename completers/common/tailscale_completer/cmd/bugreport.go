package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bugreportCmd = &cobra.Command{
	Use:   "bugreport",
	Short: "Print a shareable identifier to help diagnose issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugreportCmd).Standalone()

	bugreportCmd.Flags().Bool("diagnose", false, "run additional in-depth checks")
	bugreportCmd.Flags().Bool("record", false, "pause and then write another bugreport")
	rootCmd.AddCommand(bugreportCmd)
}
