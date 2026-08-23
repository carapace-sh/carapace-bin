package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(diagnoseCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(dumpCmd)
}

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "run diagnostic tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "display Wi-Fi information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "enable or disable logging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump Wi-Fi log buffer to file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnoseCmd).Standalone()
	carapace.Gen(infoCmd).Standalone()
	carapace.Gen(logCmd).Standalone()
	carapace.Gen(dumpCmd).Standalone()

	diagnoseCmd.Flags().StringP("output", "f", "", "Output directory path")

	carapace.Gen(diagnoseCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
	})
}