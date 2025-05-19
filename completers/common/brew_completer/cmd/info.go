package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"abv"},
	Short:   "Display brief statistics for your Homebrew installation",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("analytics", false, "List global Homebrew analytics data")
	infoCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	infoCmd.Flags().String("category", "", "Which type of analytics data to retrieve")
	infoCmd.Flags().String("days", "", "How many days of analytics data to retrieve")
	infoCmd.Flags().Bool("debug", false, "Display any debugging information")
	infoCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks")
	infoCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	infoCmd.Flags().Bool("github", false, "Open the GitHub source page for <formula> and <cask> in a browser")
	infoCmd.Flags().Bool("help", false, "Show this message")
	infoCmd.Flags().Bool("installed", false, "Print JSON of formulae that are currently installed")
	infoCmd.Flags().Bool("json", false, "Print a JSON representation")
	infoCmd.Flags().Bool("quiet", false, "Make some output more quiet")
	infoCmd.Flags().Bool("variations", false, "Include the variations hash in each formula's JSON output")
	infoCmd.Flags().Bool("verbose", false, "Show more verbose analytics data for <formula>")
	rootCmd.AddCommand(infoCmd)

	infoCmd.MarkFlagsMutuallyExclusive("cask", "formula")

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"category": carapace.ActionValues("install", "install-on-request", "build-error", "cask-install", "os-version"),
		"days":     carapace.ActionValues("30", "90", "365"),
	})

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				brew.ActionAllCasks().Unless(infoCmd.Flag("formula").Changed),
				brew.ActionAllFormulae().Unless(infoCmd.Flag("cask").Changed),
			).ToA().FilterArgs()
		}),
	)
}
