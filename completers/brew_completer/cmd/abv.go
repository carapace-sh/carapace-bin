package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var abvCmd = &cobra.Command{
	Use:   "abv",
	Short: "Display brief statistics for your Homebrew installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(abvCmd).Standalone()

	abvCmd.Flags().Bool("analytics", false, "List global Homebrew analytics data")
	abvCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	abvCmd.Flags().String("category", "", "Which type of analytics data to retrieve")
	abvCmd.Flags().String("days", "", "How many days of analytics data to retrieve")
	abvCmd.Flags().Bool("debug", false, "Display any debugging information")
	abvCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks")
	abvCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	abvCmd.Flags().Bool("github", false, "Open the GitHub source page for <formula> and <cask> in a browser")
	abvCmd.Flags().Bool("help", false, "Show this message")
	abvCmd.Flags().Bool("installed", false, "Print JSON of formulae that are currently installed")
	abvCmd.Flags().Bool("json", false, "Print a JSON representation")
	abvCmd.Flags().Bool("quiet", false, "Make some output more quiet")
	abvCmd.Flags().Bool("variations", false, "Include the variations hash in each formula's JSON output")
	abvCmd.Flags().Bool("verbose", false, "Show more verbose analytics data for <formula>")
	rootCmd.AddCommand(abvCmd)

	abvCmd.MarkFlagsMutuallyExclusive("cask", "formula")

	carapace.Gen(abvCmd).FlagCompletion(carapace.ActionMap{
		"category": carapace.ActionValues("install", "install-on-request", "build-error", "cask-install", "os-version"),
		"days":     carapace.ActionValues("30", "90", "365"),
	})

	carapace.Gen(abvCmd).PositionalAnyCompletion(
		carapace.Batch(
			brew.ActionAllCasks().Unless(func(c carapace.Context) bool { return abvCmd.Flag("formula").Changed }),
			brew.ActionAllFormulae().Unless(func(c carapace.Context) bool { return abvCmd.Flag("cask").Changed }),
		).ToA(),
	)
}
