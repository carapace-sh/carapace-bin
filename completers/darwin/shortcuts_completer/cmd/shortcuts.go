package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all shortcuts or shortcut folders",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a shortcut",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a shortcut file for sharing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Open a shortcut in the Shortcuts editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().StringP("folder", "f", "", "List shortcuts only from a specific folder")
	listCmd.Flags().Bool("folders", false, "List all folder names instead of shortcuts")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().StringSliceP("input-path", "i", nil, "Input files to pass to the shortcut")
	runCmd.Flags().StringP("output-path", "o", "", "Write the shortcut's output to a file")
	runCmd.Flags().String("output-type", "", "Force output to a specific UTI")
	rootCmd.AddCommand(runCmd)
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"input-path":  carapace.ActionFiles(),
		"output-path": carapace.ActionFiles(),
	})

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

	carapace.Gen(viewCmd).Standalone()
	rootCmd.AddCommand(viewCmd)
}
