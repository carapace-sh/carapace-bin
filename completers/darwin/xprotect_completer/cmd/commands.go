package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(helpCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "perform an update of XProtect assets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "print the currently online available update version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version of the currently installed XProtect assets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "display XProtect logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "print the current status of XProtect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "print help for a particular subcommand",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()
	carapace.Gen(checkCmd).Standalone()
	carapace.Gen(versionCmd).Standalone()
	carapace.Gen(logsCmd).Standalone()
	carapace.Gen(statusCmd).Standalone()
	carapace.Gen(helpCmd).Standalone()

	checkCmd.Flags().Bool("json", false, "Output in JSON format")
	versionCmd.Flags().Bool("json", false, "Output in JSON format")
	statusCmd.Flags().Bool("json", false, "Output in JSON format")
	updateCmd.Flags().Bool("prerelease", false, "Attempt to use a prerelease update")
}