package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "idea",
	Short: "IntelliJ IDEA CLI",
	Long:  "https://www.jetbrains.com/idea/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("wait", false, "Wait for the files to be closed before returning to the command prompt")

	// TODO this should only complete one file
	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.Batch(
		carapace.ActionFiles(),
		ActionJetbrainsPseudoFlags(),
	).ToA())
}

func ActionJetbrainsPseudoFlags() carapace.Action {
	// These behave like flags, but I don't think carapace can actually complete them as if they are
	return carapace.ActionValuesDescribed(
		"nosplash", "Do not show the splash screen when loading",
		"dontReopenProjects", "Do not reopen projects and show the welcome screen",
		"disableNonBundledPlugins", "Do not load manually installed plugins",
	).FilterArgs()
}
