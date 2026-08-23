package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spctl",
	Short: "SecAssessment system policy security",
	Long:  "https://keith.github.io/xcode-manpages/spctl.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("add", false, "Add rule(s) to the rule database")
	rootCmd.Flags().Bool("anchor", false, "Indicate arguments are hashes of anchor certificates")
	rootCmd.Flags().BoolP("assess", "a", false, "Perform an assessment on the files given")
	rootCmd.Flags().Bool("continue", false, "Continue assessing additional files after failure")
	rootCmd.Flags().Bool("disable", false, "Disable rule(s) in the rule database")
	rootCmd.Flags().Bool("disable-status", false, "Query whether the option is available")
	rootCmd.Flags().Bool("enable", false, "Enable rule(s) in the rule database")
	rootCmd.Flags().Bool("global-disable", false, "Reveal the option to allow apps from anywhere")
	rootCmd.Flags().Bool("global-enable", false, "Enable the assessment subsystem")
	rootCmd.Flags().Bool("hash", false, "Indicate arguments are code directory hashes")
	rootCmd.Flags().Bool("ignore-cache", false, "Do not query or use the assessment object cache")
	rootCmd.Flags().String("label", "", "Specify a string label to attach to or find in rules")
	rootCmd.Flags().Bool("no-cache", false, "Do not place outcomes into the assessment object cache")
	rootCmd.Flags().Bool("path", false, "Indicate arguments denote paths to files on disk")
	rootCmd.Flags().String("priority", "", "Specify the priority of the rule(s)")
	rootCmd.Flags().Bool("raw", false, "Write output as raw XML plist")
	rootCmd.Flags().Bool("remove", false, "Remove rule(s) from the rule database")
	rootCmd.Flags().Bool("requirement", false, "Indicate arguments are code requirement source")
	rootCmd.Flags().Bool("reset-default", false, "Reset the system policy database to default")
	rootCmd.Flags().Bool("rule", false, "Indicate arguments are index numbers of existing rules")
	rootCmd.Flags().Bool("status", false, "Query whether the assessment subsystem is enabled")
	rootCmd.Flags().StringP("type", "t", "", "Specify the type of assessment")
	rootCmd.Flags().BoolP("verbose", "v", false, "Request more verbose output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("execute", "install", "open"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}