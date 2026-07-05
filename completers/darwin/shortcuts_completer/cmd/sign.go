package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a shortcut file for sharing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signCmd).Standalone()
	signCmd.Flags().String("input", "", "Path to the .shortcut file to sign")
	signCmd.Flags().String("mode", "", "Signing mode: people-who-know-me or Anyone")
	signCmd.Flags().String("output", "", "Path for the signed output .shortcut file")
	rootCmd.AddCommand(signCmd)
	carapace.Gen(signCmd).FlagCompletion(carapace.ActionMap{
		"input":  carapace.ActionFiles(),
		"mode":   carapace.ActionValues("people-who-know-me", "Anyone"),
		"output": carapace.ActionFiles(),
	})
}
